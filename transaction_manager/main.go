package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"log"
	"net"
	"simple_warehouse/transaction_manager/api"
	"simple_warehouse/transaction_manager/external"
	"simple_warehouse/transaction_manager/repository/store"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

func main() {
	// Initialize database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=wms password=wms host=127.0.0.1 dbname=wms")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)
	database := store.New(conn)

	// Initialize server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterTransactionManagerServer(server, external.NewServer(database))
	log.Printf("BFF server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
