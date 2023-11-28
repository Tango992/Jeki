package repository

import "go.mongodb.org/mongo-driver/bson"

type Order interface {
	Create() (error)
	FindWithFilter(bson.M) (error)
	Update() (error)
}