package main

import (
	"flag"
	"log"

	calcpb "github.com/pieterlouw/grpc-examples/calculator/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var target = flag.String("l", ":7000", "Specify the port that the server is listening on")
var mode = flag.String("m", "sum", "Specify the client mode (sum|multiply)")

func main() {

	flag.Parse()

	ctx := context.Background()

	log.Printf("Calculator Client. Target: %v Mode: %v \n", *target, *mode)

	conn, err := grpc.Dial(*target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	calcClient := calcpb.NewCalculatorServiceClient(conn)

	if *mode == "sum" {

		rsp, err := calcClient.Sum(ctx, &calcpb.SumRequest{
			Sequence: []int64{1, 100, 33},
		})
		if err != nil {
			log.Fatalf("calcClient.Sum err: %v", err)
		}

		log.Printf("calcClient.Sum rsp: %v \n", rsp.GetResult())
	} else if *mode == "multiply" {
		rsp, err := calcClient.Multiply(ctx, &calcpb.MultiplyRequest{
			Multiplicand: 100,
			Multiplier:   12,
		})
		if err != nil {
			log.Fatalf("calcClient.Multiply err: %v", err)
		}

		log.Printf("calcClient.Multiply rsp: %v \n", rsp.GetProduct())
	} else {
		log.Fatalf("Invalid mode specified: %s", *mode)
	}

}
