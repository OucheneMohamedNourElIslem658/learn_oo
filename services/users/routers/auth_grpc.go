package routers

import (
	"context"
	"errors"

	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"
	repos "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
)

type AuthServiceServer struct {
	authpb.UnimplementedAuthServiceServer
	authRepo *repos.AuthRepository
}

func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{
		authRepo: repos.NewAuthRepository(),
	}
}

func (s *AuthServiceServer) RegisterWithEmailAndPassword(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	apiErr := s.authRepo.RegisterWithEmailAndPassword(req.FullName, req.Email, req.Password)

	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	return &authpb.RegisterResponse{
		Message:   "User registered successfully",
	}, nil
}
