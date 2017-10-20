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
	ID      string
	Address string
	Port    string
}

type args struct{}

var (
	nodes = []*serviceNode{
		{
			ID:      "backend.service1",
			Address: ":7001",
		},
		{
			ID:      "backend.service2",
			Address: ":7002",
		},
		{
			ID:      "backend.service3",
			Address: ":7003",
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

		log.Printf("Node selected: %s (%s)\n", node.ID, node.Address)

		conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
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
