package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/martian/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	sv "simple_warehouse/bff/api"
	"simple_warehouse/bff/external"
	cliPm "simple_warehouse/product_manager/api"
	cliTm "simple_warehouse/transaction_manager/api"
	"time"
)

var (
	port   = flag.Int("port", 50051, "The server_t port")
	addrPm = flag.String("addr_pm", "localhost:50052", "PM address")
	addrTm = flag.String("addr_tm", "localhost:50053", "TM address")
)

func main() {
	log.SetLevel(log.Info)

	// ************************************************************
	// Initialize the client_t
	// ************************************************************
	flag.Parse()
	// Set up a connection to the server_t.
	connPm, err := grpc.Dial(*addrPm, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("TM did not connect: %v", err)
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Errorf("PM connection did not close: %v", err)
			panic(err)
		}
	}(connPm)
	pm := cliPm.NewProductManagerClient(connPm)
	log.Infof("Product Manager client_t connected to %v", *addrPm)

	connTm, err := grpc.Dial(*addrTm, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("PM did not connect: %v", err)
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Errorf("TM connection did not close: %v", err)
			panic(err)
		}
	}(connTm)
	tm := cliTm.NewTransactionManagerClient(connTm)
	log.Infof("Transaction Manager client_t connected to %v", *addrTm)

	// Contact the server_t and print out its response.
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ************************************************************
	// Initialize the server_t
	// ************************************************************
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		panic(err)
	}
	s := grpc.NewServer()
	sv.RegisterBFFServer(s, &external.ServerT{Pm: pm, Tm: tm})
	log.Infof("BFF server_t listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
