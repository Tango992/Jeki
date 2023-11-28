package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Restaurant  Restaurant         `bson:"restaurant" json:"restaurant"`
	OrderDetail OrderDetail        `bson:"order_detail" json:"order_detail"`
	User        User               `bson:"user" json:"user"`
	Driver      Driver             `bson:"driver" json:"driver"`
	Payment     Payment            `bson:"payment" json:"payment"`
}

type Restaurant struct {
	Id      uint    `bson:"id" json:"id"`
	AdminId uint    `json:"admin_id" json:"admin_id"`
	Name    string  `bson:"name" json:"name"`
	Address Address `bson:"address" json:"address"`
	Status  string  `bson:"status" json:"status"`
}

type OrderDetail struct {
	Menus []Menu  `bson:"menus" json:"menus"`
	Total float32 `bson:"total" json:"total"`
}

type Menu struct {
	Id       uint    `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Qty      uint    `bson:"qty" json:"qty"`
	Subtotal float32 `bson:"subtotal" json:"subtotal"`
}

type User struct {
	Id      uint    `bson:"id" json:"id"`
	Name    string  `bson:"name" json:"name"`
	Address Address `bson:"address" json:"address"`
}

type Driver struct {
	Id     uint   `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Status string `bson:"status" json:"status"`
}

type Payment struct {
	Id        uint    `bson:"id" json:"id"`
	InvoiceId string  `bson:"invoice_id" json:"invoice_id"`
	Total     float32 `bson:"total" json:"total"`
	Method    string  `bson:"method" json:"method"`
	Status    string  `bson:"status" json:"status"`
}

type Address struct {
	Latitude  float32 `bson:"latitude" json:"latitude"`
	Longitude float32 `bson:"longitude" json:"longitude"`
}