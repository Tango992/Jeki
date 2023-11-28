package repository

import (
	"user-service/dto"
	"user-service/models"
)

type User interface {
	GetUserData(string) (dto.UserJoinedData, error)
	CreateUser(*models.User) (error)
}