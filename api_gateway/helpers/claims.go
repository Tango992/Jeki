package helpers

import (
	"api-gateway/dto"
	"api-gateway/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (dto.Claims, error) {
	claimsTmp := c.Get("user")
	if claimsTmp == nil {
		return dto.Claims{}, echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Failed to fetch user claims from JWT"))
	}

	claims := claimsTmp.(jwt.MapClaims)
	return dto.Claims{
		ID:    uint(claims["id"].(float64)),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		Role:  claims["role"].(string),
	}, nil
}
