package main

import (
	"log"
	"net"

	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
	"google.golang.org/grpc"

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

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middlewares.AuthorizationUnaryInterceptor()),
	)

	authService := usersRouters.NewAuthServiceServer()
	authpb.RegisterAuthServiceServer(grpcServer, authService)

	profilesService := usersRouters.NewProfilesServiceServer()
	authpb.RegisterProfilesServiceServer(grpcServer, profilesService)

	log.Printf("Listening and serving (grpc) at %v\n", "tcp:"+server.address)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
