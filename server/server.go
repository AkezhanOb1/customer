package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	"github.com/AkezhanOb1/customer/service"
)

func main() {
	address := "0.0.0.0:50055"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	pb.RegisterCustomerServiceServer(s, &service.Customer{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
