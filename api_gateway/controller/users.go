package controller

import (
	"api-gateway/dto"
	"api-gateway/models"
	"api-gateway/pb"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)
type RegisterController struct {
	Client pb.UserClient
}

func NewRegisterController(client pb.UserClient) RegisterController {
	return RegisterController{
		Client: client,
	}
}

/* 
	ROLE ID
	Customer/user = 1
	Driver = 2
	Admin = 3
*/

func (r RegisterController) RegisterUser(c echo.Context) error {
	register := new(dto.UserRegister)
	if err := c.Bind(register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Register")
	}

	if err := c.Validate(register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid validate")
	}
	
	registerData := &pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		BirthDate: register.BirthDate,
		RoleId: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	responseGrpc, err := r.Client.Register(ctx, registerData)
	if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error - failed user register")
    }

	responseData := models.User{
		ID: responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		BirthDate: register.BirthDate,
		Role: "user",
		CreatedAt: responseGrpc.CreatedAt,
	}

	response := dto.Response{
        Message: "User Berhasil register",
        Data:    responseData,
    }

    return c.JSON(http.StatusCreated, response)
}


func (r RegisterController) RegisterDriver(c echo.Context) error{
	register := new(dto.UserRegister)
	if err := c.Bind(register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Register")
	}

	if err := c.Validate(register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid validate")
	}
	registerData := &pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		BirthDate: register.BirthDate,
		RoleId: 2,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	responseGrpc, err := r.Client.Register(ctx, registerData)
	if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error - failed user register")
    }

	responseData := models.User{
		ID: responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		BirthDate: register.BirthDate,
		Role: "driver",
		CreatedAt: responseGrpc.CreatedAt,
	}

	response := dto.Response{
        Message: "User Berhasil register",
        Data:    responseData,
    }

    return c.JSON(http.StatusCreated, response)
}

func (r RegisterController) RegisterAdmin(c echo.Context) error {
	register := new(dto.UserRegister)
	if err := c.Bind(register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Register")
	}

	if err := c.Validate(register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid validate")
	}

	registerData := &pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		BirthDate: register.BirthDate,
		RoleId: 3,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	responseGrpc, err := r.Client.Register(ctx, registerData)
	if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error - failed user register")
    }

	responseData := models.User{
		ID: responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		BirthDate: register.BirthDate,
		Role: "admin",
		CreatedAt: responseGrpc.CreatedAt,
	}

	response := dto.Response{
        Message: "Admin Berhasil register",
        Data:    responseData,
    }
    return c.JSON(http.StatusCreated, response)
}

// Login
