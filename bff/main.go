package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	sv "simple_warehouse/bff/api"
	"simple_warehouse/bff/external"
	cli "simple_warehouse/product_manager/api"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server_t port")
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
)

func main() {
	// ************************************************************
	// Initialize the client_t
	// ************************************************************
	flag.Parse()
	// Set up a connection to the server_t.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Client connection did not close: %v", err)
		}
	}(conn)

	client := cli.NewProductManagerClient(conn)
	log.Printf("Product Manager client_t connected to %v", *addr)

	// Contact the server_t and print out its response.
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ************************************************************
	// Initialize the server_t
	// ************************************************************
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sv.RegisterBFFServer(s, &external.serverT{client: client})
	log.Printf("BFF server_t listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
