package routers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	authpb "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/grpc"
	authRepos "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/repositories"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServiceServer struct {
	authpb.UnimplementedAuthServiceServer
	authRepo  *authRepos.AuthRepository
	usersRepo *authRepos.ProfilesRepository
}

func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{
		authRepo: authRepos.NewAuthRepository(),
		usersRepo: authRepos.NewProfilesRepository(),
	}
}

func (s *AuthServiceServer) RegisterWithEmailAndPassword(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	apiErr := s.authRepo.RegisterWithEmailAndPassword(*req.FullName, *req.Email, *req.Password)

	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	msg := "your account has been created, please verify your email to have full access"
	return &authpb.RegisterResponse{
		Message: &msg,
	}, nil
}

func (s *AuthServiceServer) LoginWithEmailAndPassword(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	idToken, refreshToken, apiErr := s.authRepo.LoginWithEmailAndPassword(*req.Email, *req.Password)

	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	return &authpb.LoginResponse{
		IdToken:      idToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceServer) SendEmailVerificationLink(ctx context.Context, req *authpb.EmailLinkRequest) (*emptypb.Empty, error) {
	link := getDomainLink(ctx) + "/api/v1/users/auth/serve-email-verification-template"

	apiErr := s.authRepo.SendEmailVerificationLink(*req.Email, link)

	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	return &emptypb.Empty{}, nil
}

func (s *AuthServiceServer) SendPasswordResetLink(ctx context.Context, req *authpb.EmailLinkRequest) (*emptypb.Empty, error) {
	link := getDomainLink(ctx) + "/api/v1/users/auth/serve-reset-password-form"

	apiErr := s.authRepo.SendPasswordResetLink(*req.Email, link)

	if apiErr != nil {
		return nil, errors.New(apiErr.Message)
	}

	return &emptypb.Empty{}, nil
}

func (s *AuthServiceServer) RefreshIDToken(ctx context.Context, req *emptypb.Empty) (*authpb.RefreshIDTokenReponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata not found in context")
	}

	if tokens := md["authorization"]; len(tokens) > 0 {
		idToken, apiErr := s.authRepo.RefreshIdToken(md["authorization"][0])
		if apiErr != nil {
			return nil, errors.New(apiErr.Message)
		}

		return &authpb.RefreshIDTokenReponse{
			IdToken: idToken,
		}, nil
	}

	return nil, errors.New("invalid authorization")
}

func getDomainLink(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("metadata not found in context")
	}

	var domain string

	if authorities := md[":authority"]; len(authorities) > 0 {
		domain = authorities[0]
	} else if hosts := md["host"]; len(hosts) > 0 {
		domain = hosts[0]
	} else {
		fmt.Println("no authority or host found in metadata")
	}

	domain = strings.Split(domain, ":")[0]
	link := fmt.Sprintf("http://%s", domain+":8000")

	return link
}
