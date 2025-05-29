package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"
)

func runTestGRPC() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.RegisterWithEmailAndPassword(ctx, &authpb.RegisterRequest{
		FullName: "John Doe",
		Email:    "john@example.com",
		Password: "securepassword123",
	})

	if err != nil {
		log.Fatalf("gRPC call failed: %v", err)
	}

	log.Printf("Response: %s", resp.Message)
}
