package repository

import (
	"user-service/dto"
	"user-service/models"
)

type User interface {
	GetUserData(string) (dto.UserJoinedData, error)
	CreateUser(*models.User) error
	CreateDriverData(userID uint32) error
	AddToken(*models.Verification) error
	GetAvailableDriver() (dto.DriverData, error)
	SetDriverStatusOnline(driverID uint) error
	SetDriverStatusOngoing(driverID uint) error
	SetDriverStatusOffline(driverID uint) error
	VerifyNewUser(id uint32, token string) error
}
