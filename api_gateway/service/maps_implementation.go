package service

import (
	"api-gateway/dto"
	"api-gateway/utils"
	"context"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

type GoogleMaps struct {
	Client *maps.Client
}

func NewMapsService() Maps {
	apiKey := os.Getenv("MAPS_API_KEY")
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	
	return GoogleMaps{
		Client: client,
	}
}

// Returns latitude and longitude from the given address
func (g GoogleMaps) GetCoordinate(address string) (dto.Coordinate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	request := &maps.GeocodingRequest{
		Address: address,
	}
	
	results, err := g.Client.Geocode(ctx, request)
	if err != nil {
		return dto.Coordinate{}, echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	if len(results) < 1 {
		return dto.Coordinate{}, echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails("Invalid address"))
	}
	
	result := results[0]
	coordinate := dto.Coordinate{
		Latitude: float32(result.Geometry.Location.Lat),
		Longitude: float32(result.Geometry.Location.Lng),
	}
	return coordinate, nil
}