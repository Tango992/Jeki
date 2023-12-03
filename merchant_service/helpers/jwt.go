package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SignJwtForGrpc() (string, error) {
	secret := os.Getenv("SERVICES_JWT_SECRET")
	
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
	}
	
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", status.Error(codes.Internal, err.Error())
    }
    return tokenString, nil
}