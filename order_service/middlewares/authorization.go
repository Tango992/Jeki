package middlewares

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
)

func JWTAuth(ctx context.Context) (context.Context, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		fmt.Println(tokenString)
		return nil, fmt.Errorf("Couldn't get token : %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Couldn't get token : %v", err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return ctx, nil
	}
	return nil, fmt.Errorf("Failed to verify JWT")
}
