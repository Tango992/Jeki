package middlewares

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func JWTAuth(ctx context.Context) (context.Context, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		fmt.Println(tokenString)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SERVICES_JWT_SECRET")), nil
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return ctx, nil
	}
	return nil, status.Error(codes.Unauthenticated, "failed to verify jwt")
}
