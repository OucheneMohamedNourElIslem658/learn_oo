package main

import (
	"context"
	"log"
	// "time"

	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Make sure the server is running and implements ProfilesService as defined in your proto file.
// Also, ensure the proto package and service names match between client and server.

func runTestGRPC() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := authpb.NewProfilesServiceClient(conn)

	ctx := context.Background()
	// defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JfaWQiOiJmOWU3ZjBhNy1jYTZjLTQ3ZDEtYTE5Ni01YjY3MzRmYTNlYjkiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZXhwIjoxNzQ4Nzg1Mzk2LCJpZCI6IjM3OGYyMDE0LTUzMzEtNDJlYy1hZjUwLTViZDg5Y2IwMzc2NyJ9.BsnM9Fwng7zW0PwHP5D1TUkDeY3BXwxWckWorMbbUf0")

	resp, err := client.UpgradeToAuthor(ctx, nil)

	if err != nil {
		log.Fatalf("gRPC call failed: %v", err)
	}

	log.Printf("Response: %s", resp)
}
