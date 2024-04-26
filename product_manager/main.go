package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/martian/log"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"net"
	"simple_warehouse/product_manager/api"
	"simple_warehouse/product_manager/external"
	"simple_warehouse/product_manager/repository/store"
)

var (
	port = flag.Int("port", 50052, "The server port")
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
	api.RegisterProductManagerServer(server, external.NewServer(database))
	log.Infof("PM server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
