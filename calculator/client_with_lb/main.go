package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	calcpb "github.com/pieterlouw/grpc-examples/calculator/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type serviceNode struct {
	id      string
	address string
}

type args struct{}

var (
	nodes = []*serviceNode{
		{
			id:      "backend.service1",
			address: ":7001",
		},
		{
			id:      "backend.service2",
			address: ":7002",
		},
		{
			id:      "backend.service3",
			address: ":7003",
		},
	}
	c uint64
)

func main() {

	ops := []struct {
		name string
		args []int64
	}{
		{"sum", []int64{1, 100, 33}},
		{"sum", []int64{2, 100, 33}},
		{"sum", []int64{3, 100, 33}},
		{"multiply", []int64{4, 100}},
		{"multiply", []int64{5, 100}},
	}

	//flag.Parse()

	ctx := context.Background()

	for _, op := range ops {

		// get service node to call
		node, err := selectNode()
		if err != nil {
			log.Fatalf("err selectNode %v", err)
		}

		log.Printf("Node selected: %s (%s)\n", node.id, node.address)

		conn, err := grpc.Dial(node.address, grpc.WithInsecure()) // <-- One can implement a connection pool and select a connection from the pool here instead of creating a new one
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}

		calcClient := calcpb.NewCalculatorServiceClient(conn)

		if op.name == "sum" {

			rsp, err := calcClient.Sum(ctx, &calcpb.SumRequest{
				Sequence: op.args,
			})
			if err != nil {
				log.Fatalf("calcClient.Sum err: %v", err)
			}

			log.Printf("calcClient.Sum rsp: %v \n", rsp.GetResult())
		} else if op.name == "multiply" {
			rsp, err := calcClient.Multiply(ctx, &calcpb.MultiplyRequest{
				Multiplicand: op.args[0],
				Multiplier:   op.args[1],
			})
			if err != nil {
				log.Fatalf("calcClient.Multiply err: %v", err)
			}

			log.Printf("calcClient.Multiply rsp: %v \n", rsp.GetProduct())
		} else {
			log.Fatalf("Invalid mode specified: %s", op.name)
		}
	}

}

// selectNode performs round robin on nodes
func selectNode() (*serviceNode, error) {

	var mtx sync.Mutex

	if len(nodes) == 0 {
		return nil, fmt.Errorf("No nodes specified")
	}

	mtx.Lock()

	old := atomic.AddUint64(&c, 1) - 1

	node := nodes[old%uint64(len(nodes))]
	mtx.Unlock()

	return node, nil

}
