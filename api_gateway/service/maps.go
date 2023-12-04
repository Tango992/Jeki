package service

import "api-gateway/dto"

type Maps interface {
	GetCoordinate(string) (dto.Coordinate, error)
}