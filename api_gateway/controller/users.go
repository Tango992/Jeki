package controller

import (
	"api-gateway/dto"
	"api-gateway/helpers"
	"api-gateway/models"
	"api-gateway/pb"
	"api-gateway/utils"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Client pb.UserClient
}

func NewUserController(client pb.UserClient) UserController {
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

	registerData := &pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
		BirthDate: register.BirthDate,
		RoleId:    1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	responseGrpc, err := u.Client.Register(ctx, registerData)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	responseData := models.User{
		ID:        responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  "-",
		BirthDate: register.BirthDate,
		Role:      "user",
		CreatedAt: responseGrpc.CreatedAt,
	}

	response := dto.Response{
		Message: "Registered succesfully",
		Data:    responseData,
	}

	return c.JSON(http.StatusCreated, response)
}

func (u UserController) RegisterDriver(c echo.Context) error {
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

	registerData := &pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
		BirthDate: register.BirthDate,
		RoleId:    2,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	responseGrpc, err := u.Client.Register(ctx, registerData)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	responseData := models.User{
		ID:        responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  "-",
		BirthDate: register.BirthDate,
		Role:      "driver",
		CreatedAt: responseGrpc.CreatedAt,
	}

	response := dto.Response{
		Message: "Registered succesfully",
		Data:    responseData,
	}

	return c.JSON(http.StatusCreated, response)
}

func (u UserController) RegisterAdmin(c echo.Context) error {
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

	registerData := &pb.RegisterRequest{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
		BirthDate: register.BirthDate,
		RoleId:    3,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	responseGrpc, err := u.Client.Register(ctx, registerData)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	responseData := models.User{
		ID:        responseGrpc.UserId,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  "-",
		BirthDate: register.BirthDate,
		Role:      "admin",
		CreatedAt: responseGrpc.CreatedAt,
	}

	response := dto.Response{
		Message: "Registered succesfully",
		Data:    responseData,
	}
	return c.JSON(http.StatusCreated, response)
}

// Login
func (u UserController) LoginUser(c echo.Context) error {
	loginReq := new(dto.UserLogin)
	if err := c.Bind(loginReq); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(loginReq); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	emailRequest := &pb.EmailRequest{
		Email: loginReq.Email,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userDataTmp, err := u.Client.GetUserData(ctx, emailRequest)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	userData := models.User{
		ID: userDataTmp.Id,
		FirstName: userDataTmp.FirstName,
		LastName: userDataTmp.LastName,
		Email: userDataTmp.Email,
		Password: userDataTmp.Password,
		BirthDate: userDataTmp.BirthDate,
		Role: userDataTmp.Role,
	}

	// check hashpassword
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginReq.Password))
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	// jika password match, kirim token
	

	// handle jwt dan error, get generate token
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	response := dto.Response{
		Message: "Login succesfully",
		Data:   token,
	}
	return c.JSON(http.StatusOK, response)

}
