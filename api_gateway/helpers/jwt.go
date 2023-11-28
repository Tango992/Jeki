package helpers

import (
	"fmt"
	"net/http"
	"os"
	"api-gateway/models"
	"api-gateway/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func SignNewJWT(c echo.Context, user models.User) error{
	claims := jwt.MapClaims{
		"exp": time.Now().Add(4 * time.Hour).Unix(),
		"id": user.ID,
		"email": user.Email,
		"name": user.FirstName,
		"role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = tokenString
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Expires = time.Now().Add(2 * time.Hour)
	c.SetCookie(cookie)

	return nil
}

func SignJwtForGrpc(secret string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
	}
	
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", fmt.Errorf("failed to generate JWT: %v", err)
    }
    return tokenString, nil
}