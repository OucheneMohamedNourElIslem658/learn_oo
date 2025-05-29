package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"

	usersRouters "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/routers"
)

type GRPCServer struct {
	address string
}

func NewGRPCServer(address string) *GRPCServer {
	return &GRPCServer{
		address: address,
	}
}

func (server *GRPCServer) Run() {
	lis, err := net.Listen("tcp", server.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authService := usersRouters.NewAuthServiceServer()

	authpb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Listening and serving (grpc) at %v\n", "tcp:"+server.address)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
