package main

import (
	"flag"
	"log"
	"net"

	context "golang.org/x/net/context"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
)

var (
	listenPort = flag.String("l", ":7000", "Specify the port that the server will listen on")
	proxyDest  = flag.String("d", ":7001", "Specify the port that the server will proxy to")
	mode       = flag.String("m", "transparent", "Specify the mode of proxying (transparent|explicit)")
)

func main() {
	flag.Parse()

	director := func(ctx context.Context, fullMethodName string) (*grpc.ClientConn, error) {
		log.Println("director: ", fullMethodName)

		return grpc.DialContext(ctx, *proxyDest, grpc.WithInsecure(), grpc.WithCodec(proxy.Codec()))
	}

	var server *grpc.Server
	if *mode == "transparent" {
		// A gRPC server with the proxying codec enabled.
		server = grpc.NewServer(grpc.CustomCodec(proxy.Codec()),
			grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))

	} else if *mode == "explicit" {

		server = grpc.NewServer(grpc.CustomCodec(proxy.Codec()))
		// Register a Service with  it's methods explicitly.
		proxy.RegisterService(server, director, "proto.CalculatorService", "Sum")
	} else {
		log.Fatalf("invalid mode: %v", *mode)
	}

	lis, err := net.Listen("tcp", *listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", *listenPort, ". Proxying to ", *proxyDest, " in ", *mode, " mode")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
