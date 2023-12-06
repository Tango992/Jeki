package helpers

import (
	"math"
	"order-service/model"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
)

// Returns distance in kilometer
func calculateDistance(restaurant model.Restaurant, user model.User) (float64) {
	userLocation := orb.Point{
		float64(user.Address.Longitude), 
		float64(user.Address.Latitude),
	}
	restaurantLocation := orb.Point{
		float64(restaurant.Address.Longitude),
		float64(restaurant.Address.Latitude),
	}
	distance := geo.Distance(userLocation, restaurantLocation)
	return distance / 1000
}

func CalculateDeliveryFee(restaurant model.Restaurant, user model.User) float32 {
	var (
		platformFee float32 = 4500
		pricePerKm float32 = 2700
	)

	distance := calculateDistance(restaurant, user)
	deliveryFee := float32(distance) * pricePerKm
	totalFee := deliveryFee + platformFee
	
	if totalFee < 10000 {
		return 10000
	}
	return float32(math.Round(float64(totalFee)))
}
