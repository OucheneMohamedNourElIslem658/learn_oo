package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	ContextKeyID       contextKey = "id"
	ContextKeyAuthorID contextKey = "author_id"
)

func AuthorizationUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("metadata not found in context")
		}

		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			return handler(ctx, req)
		}

		authorization := authHeaders[0]
		if strings.HasPrefix(authorization, "Bearer ") {
			idToken := strings.TrimPrefix(authorization, "Bearer ")
			if idToken != "" {
				claims, isValid, err := utils.VerifyIDToken(idToken)
				if err != nil {
					return nil, errors.New("id token expired")
				}
				if !isValid {
					return nil, errors.New("invalid id token")
				}
				ctx = context.WithValue(ctx, "id", claims.ID)
				if claims.AuthorID != nil {
					ctx = context.WithValue(ctx, "author_id", *claims.AuthorID)
				}
			}
		}
		return handler(ctx, req)
	}
}