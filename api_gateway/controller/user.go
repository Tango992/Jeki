package controller

import (
	"api-gateway/dto"
	"api-gateway/helpers"
	"api-gateway/models"
	userpb "api-gateway/pb/userpb"
	"api-gateway/utils"
	"net/http"
	"strconv"

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

// @Summary     Register a new user
// @Description Register a new user with the role 'User'
// @Tags        customer
// @Accept      json
// @Produce     json
// @Param       request body dto.UserRegister true "User registration details"
// @Success     201 {object} dto.SwaggerResponseRegister
// @Failure     400 {object} utils.ErrResponse
// @Failure     409 {object} utils.ErrResponse
// @Failure     500 {object} utils.ErrResponse
// @Router      /users/register/user [post]
func (u UserController) RegisterUser(c echo.Context) error {
	return u.Register(c, userRoleID, userRole)
}

// @Summary 	Register a new driver
// @Description Register a new user with the role 'Driver'
// @Tags		driver
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.UserRegister true "Driver registration details"
// @Success 	201 {object} dto.SwaggerResponseRegister
// @Failure 	400 {object} utils.ErrResponse
// @Failure     409 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/register/driver [post]
func (u UserController) RegisterDriver(c echo.Context) error {
	return u.Register(c, driverRoleID, driverRole)
}

// @Summary 	Register a new admin
// @Description Register a new user with the role 'Admin'
// @Tags		merchant
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.UserRegister true "Admin registration details"
// @Success 	201 {object} dto.SwaggerResponseRegister
// @Failure 	400 {object} utils.ErrResponse
// @Failure     409 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/register/admin [post]
func (u UserController) RegisterAdmin(c echo.Context) error {
	return u.Register(c, adminRoleID, adminRole)
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

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
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

	if roleName == driverRole {
		ctx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		if _, err := u.Client.CreateDriverData(ctx, &userpb.DriverId{Id: responseGrpc.UserId}); err != nil {
			return helpers.AssertGrpcStatus(err)
		}
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

// @Summary 	User login
// @Description Authenticate and login a user
// @Tags        all user
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.UserLogin true "User login details"
// @Success     200  {object}  dto.Response
// @Failure     400  {object}  utils.ErrResponse
// @Failure     401  {object}  utils.ErrResponse
// @Failure     500  {object}  utils.ErrResponse
// @Router      /users/users/login [post]
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

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	userDataTmp, err := u.Client.GetUserData(ctx, emailRequest)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid username/password"))
			default:
				return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(e.Message()))
			}
		}
	}

	if !userDataTmp.Verified {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Please do an email verification first"))
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

	if userData.Role == driverRole {
		ctx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		if _, err := u.Client.SetDriverStatusOnline(ctx, &userpb.DriverId{Id: userData.ID}); err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
		}
	}

	if err := helpers.SignNewJWT(c, userData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Login succesfully",
		Data:    "Authorization is stored in cookie",
	})
}

// @Summary 	Logout the user
// @Description Logout the currently authenticated user
// @Tags        all user
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.Response
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/logout [get]
func (u UserController) Logout(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	userIdPb := &userpb.DriverId{
		Id: uint32(user.ID),
	}

	if user.Role == driverRole {
		ctx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		if _, err := u.Client.SetDriverStatusOffline(ctx, userIdPb); err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
		}
	}

	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = ""
	cookie.SameSite = http.SameSiteLaxMode
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Logged out",
		Data:    "Authorization cookie has been deleted",
	})
}

// @Summary 	Verify user registration
// @Description Verify the user registration using token sent through an email
// @Tags        all user
// @Accept 		json
// @Produce 	json
// @Param 		userid path integer true "User ID"
// @Param 		token path string true "Verification token"
// @Success 	200 {object} dto.Response
// @Failure 	400 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/verify/{userid}/{token} [get]
func (u UserController) VerifyUser(c echo.Context) error {
	token := c.Param("token")
	userIdTmp := c.Param("userid")
	userId, err := strconv.Atoi(userIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails("Invalid verification URL"))
	}

	pbUserData := &userpb.UserCredential{
		Id:    uint32(userId),
		Token: token,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	if _, err := u.Client.VerifyNewUser(ctx, pbUserData); err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/user/verified")
}
