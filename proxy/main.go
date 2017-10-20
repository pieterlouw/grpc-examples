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
	// listen on one port

	// proxy per service + method

	// proxy per service

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

	//pb_test.RegisterTestServiceServer(server, &testImpl{})

}

/*
	customDirector := func(ctx context.Context, fullMethodName string) (*grpc.ClientConn, error) {
		// Make sure we never forward internal services.
		if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
			return nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
		}
		md, ok := metadata.FromContext(ctx)
		if ok {
			// Decide on which backend to dial
			if val, exists := md[":authority"]; exists && val[0] == "staging.api.example.com" {
				// Make sure we use DialContext so the dialing can be cancelled/time out together with the context.
				return grpc.DialContext(ctx, "api-service.staging.svc.local", grpc.WithCodec(proxy.Codec()))
			} else if val, exists := md[":authority"]; exists && val[0] == "api.example.com" {
				return grpc.DialContext(ctx, "api-service.prod.svc.local", grpc.WithCodec(proxy.Codec()))
			}
		}
		return nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}

	server := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(customDirector)))
*/

//backendConn, err := grpc.Dial(*proxyDest, grpc.WithInsecure())
//if err != nil {
//	log.Fatalf("grpc.Dial err: %v", err)
//}

//var director proxy.StreamDirector
