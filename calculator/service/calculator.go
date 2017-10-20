package main

import (
	"flag"
	"log"
	"net"

	calcpb "github.com/pieterlouw/grpc-examples/calculator/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var listenPort = flag.String("l", ":7001", "Specify the port that the server will listen on")

type calculatorService struct{}

func (c *calculatorService) Sum(ctx context.Context, r *calcpb.SumRequest) (*calcpb.SumResponse, error) {

	log.Printf("Sum Request: %+v\n", r.GetSequence())

	sequence := r.GetSequence()
	var sumResult int64

	for _, value := range sequence {
		sumResult += value
	}

	return &calcpb.SumResponse{
		Result: sumResult,
	}, nil
}
func (c *calculatorService) Multiply(ctx context.Context, req *calcpb.MultiplyRequest) (*calcpb.MultiplyResponse, error) {
	var product int64

	log.Printf("Multiply Request: %d * %d\n", req.GetMultiplicand(), req.GetMultiplier())

	product = req.GetMultiplicand() * req.GetMultiplier()

	return &calcpb.MultiplyResponse{
		Product: product,
	}, nil
}

func main() {
	flag.Parse()
	calcService := &calculatorService{}

	lis, err := net.Listen("tcp", *listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", *listenPort)
	server := grpc.NewServer()

	calcpb.RegisterCalculatorServiceServer(server, calcService)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
