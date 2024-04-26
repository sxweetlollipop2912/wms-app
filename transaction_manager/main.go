package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/martian/log"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"net"
	"simple_warehouse/transaction_manager/api"
	"simple_warehouse/transaction_manager/external"
	"simple_warehouse/transaction_manager/repository/store"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

func main() {
	log.SetLevel(log.Info)

	// Initialize database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=wms password=wms host=127.0.0.1 dbname=wms")
	if err != nil {
		log.Errorf("Unable to connect to database: %v", err)
		panic(err)
	}
	defer conn.Close(ctx)
	database := store.New(conn)

	// Initialize server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		panic(err)
	}
	server := grpc.NewServer()
	api.RegisterTransactionManagerServer(server, external.NewServer(database))
	log.Infof("TM server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
