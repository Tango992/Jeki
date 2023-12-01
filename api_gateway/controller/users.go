package controller

import (
	"api-gateway/dto"
	"api-gateway/helpers"
	"api-gateway/models"
	userpb "api-gateway/pb/userpb"
	"api-gateway/utils"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	Client userpb.UserClient
}

func NewUserController(client userpb.UserClient) UserController {
	return UserController{
		Client: client,
	}
}

/*
	ROLE ID
	Customer/user = 1
	Driver = 2
	Admin = 3
*/

func (u UserController) RegisterUser(c echo.Context) error {
	return u.Register(c, 1, "user")
}

func (u UserController) RegisterDriver(c echo.Context) error {
	return u.Register(c, 2, "driver")
}

func (u UserController) RegisterAdmin(c echo.Context) error {
	return u.Register(c, 3, "admin")
}

func (u UserController) Register(c echo.Context, roleId uint, roleName string) error {
	register := new(dto.UserRegister)
	if err := c.Bind(register); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(register); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := helpers.DateValidator(register.BirthDate); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	registerData := &userpb.RegisterRequest{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
		BirthDate: register.BirthDate,
		RoleId:    uint32(roleId),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	responseGrpc, err := u.Client.Register(ctx, registerData)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.AlreadyExists:
				return echo.NewHTTPError(utils.ErrConflict.EchoFormatDetails(e.Message()))
			case codes.Internal:
				return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(e.Message()))
			}
		}
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	responseData := models.User{
		ID:        responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		BirthDate: register.BirthDate,
		Role:      roleName,
		CreatedAt: responseGrpc.CreatedAt,
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Registered succesfully",
		Data:    responseData,
	})
}

func (u UserController) Login(c echo.Context) error {
	loginReq := new(dto.UserLogin)
	if err := c.Bind(loginReq); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(loginReq); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	emailRequest := &userpb.EmailRequest{
		Email: loginReq.Email,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userDataTmp, err := u.Client.GetUserData(ctx, emailRequest)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid username/password"))
			case codes.Internal:
				return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(e.Message()))
			}
		}
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	userData := models.User{
		ID:        userDataTmp.Id,
		FirstName: userDataTmp.FirstName,
		LastName:  userDataTmp.LastName,
		Email:     userDataTmp.Email,
		Password:  userDataTmp.Password,
		BirthDate: userDataTmp.BirthDate,
		Role:      userDataTmp.Role,
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginReq.Password)); err != nil {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid username/password"))
	}

	if err := helpers.SignNewJWT(c, userData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Login succesfully",
		Data:    "Authorization is stored in cookie",
	})
}
