package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	cli "simple_warehouse/bff/api"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
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

	client := cli.NewBFFClient(conn)
	// Contact the server_t and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	importRequests := [2]cli.ImportRequest{
		{
			Sku:             "abc1",
			Name:            "name1",
			ExpiredDate:     1e9,
			Category:        "category1",
			ShelfNames:      []string{"shelf1", "shelf2"},
			QuantityOnShelf: []int64{1, 2},
		},
		{
			Sku:             "abc2",
			Name:            "name2",
			ExpiredDate:     1e9,
			Category:        "category2",
			ShelfNames:      []string{"shelf1", "shelf2"},
			QuantityOnShelf: []int64{1, 2},
		},
	}

	getRequests := [2]cli.GetProductRequest{
		{
			Sku: "abc1",
		},
		{
			Sku: "abc2",
		},
	}

	for _, req := range importRequests {
		r, err := client.Import(ctx, &req)
		if err != nil || !r.Ok {
			log.Fatalf("Import failed: %v", err)
		}
	}

	for _, req := range getRequests {
		r, err := client.GetProduct(ctx, &req)
		if err != nil || !r.Ok {
			log.Fatalf("GetProduct failed: %v", err)
		}
		log.Printf("GetProduct: %v", r)
	}

	defer cancel()
}
