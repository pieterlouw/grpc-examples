package main

import (
	"flag"
	"log"
	"net"

	stringspb "github.com/pieterlouw/grpc-examples/strings/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var listenPort = flag.String("l", ":7002", "Specify the port that the server will listen on")

type stringsService struct{}

func (c *stringsService) Reverse(ctx context.Context, r *stringspb.ReverseRequest) (*stringspb.ReverseResponse, error) {
	s := r.GetValue()
	val := []rune(s[:])

	for i, j := 0, len(val)-1; i < len(val)/2; i, j = i+1, j-1 {
		val[i], val[j] = val[j], val[i]
	}

	return &stringspb.ReverseResponse{
		Reversed: string(val),
	}, nil
}

func (c *stringsService) Concat(ctx context.Context, r *stringspb.ConcatRequest) (*stringspb.ConcatResponse, error) {
	return &stringspb.ConcatResponse{
		Concatenated: r.GetValue1() + r.GetValue2(),
	}, nil
}

func main() {
	flag.Parse()
	strService := &stringsService{}

	lis, err := net.Listen("tcp", *listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", *listenPort)
	server := grpc.NewServer()

	stringspb.RegisterStringsServiceServer(server, strService)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
