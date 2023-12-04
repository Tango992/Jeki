package controller

import (
	"api-gateway/dto"
	"api-gateway/helpers"
	"api-gateway/pb/orderpb"
	"api-gateway/service"
	"api-gateway/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Client orderpb.OrderServiceClient
	Maps service.Maps
}

func NewOrderController(client orderpb.OrderServiceClient, maps service.Maps) OrderController {
	return OrderController{
		Client: client,
		Maps: maps,
	}
}

func (o OrderController) UserCreateOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != userRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}

	var orderRequest dto.NewOrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&orderRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	coordinate, err := o.Maps.GetCoordinate(orderRequest.Address)
	if err != nil {
		return err
	}

	pbOrderItems := helpers.AssertToPbOrderItems(orderRequest)
	pbOrderRequest := &orderpb.RequestOrderData{
		UserId: uint32(user.ID),
		Name: user.Name,
		Email: user.Email,
		Address: &orderpb.Address{
			Latitude: coordinate.Latitude,
			Longitude: coordinate.Longitude,
		},
		OrderItems: pbOrderItems,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	response, err := o.Client.PostOrder(ctx, pbOrderRequest)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Order successfully created",
		Data: response,
	})
}

func (o OrderController) UsersGetAllOrders(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != userRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.Client.GetUserAllOrders(ctx, &orderpb.UserId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all user's orders",
		Data: orders,
	})
}

func (o OrderController) UsersGetOngoingOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != userRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.Client.GetUserCurrentOrders(ctx, &orderpb.UserId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get user's ongoing orders",
		Data: orders,
	})
}

func (o OrderController) GetOrderById(c echo.Context) error {
	orderId := c.Param("id")
	
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != userRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	order, err := o.Client.GetOrderById(ctx, &orderpb.OrderId{Id: orderId})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get order by ID",
		Data: order,
	})
}

func (o OrderController) MerchantGetAllOrders(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.Client.GetRestaurantAllOrders(ctx, &orderpb.AdminId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all merchant's orders",
		Data: orders,
	})
}

func (o OrderController) MerchantGetOngoingOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.Client.GetRestaurantCurrentOrders(ctx, &orderpb.AdminId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get merchant's current orders",
		Data: orders,
	})
}

func (o OrderController) MerchantUpdateOrder(c echo.Context) error {
	orderId := c.Param("id")

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only user role is allowed"))
	}

	var statusRequest dto.UpdateOrderStatus
	if err := c.Bind(&statusRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	if err := c.Validate(&statusRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	pbRequestData := &orderpb.RequestUpdateData{
		UserId: uint32(user.ID),
		OrderId: orderId,
		Status: statusRequest.Status,
	}

	_, err = o.Client.UpdateRestaurantOrderStatus(ctx, pbRequestData)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Order status updated",
		Data: pbRequestData,
	})
}

func (o OrderController) DriverGetAllOrders(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != driverRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only driver role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.Client.GetDriverAllOrders(ctx, &orderpb.DriverId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all driver's orders",
		Data: orders,
	})
}

func (o OrderController) DriverGetCurrentOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != driverRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only driver role is allowed"))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	order, err := o.Client.GetDriverCurrentOrder(ctx, &orderpb.DriverId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get driver's current order",
		Data: order,
	})
}

func (o OrderController) DriverUpdateOrder(c echo.Context) error {
	orderId := c.Param("id")

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != driverRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only driver role is allowed"))
	}

	var statusRequest dto.UpdateOrderStatus
	if err := c.Bind(&statusRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	if err := c.Validate(&statusRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	pbRequestData := &orderpb.RequestUpdateData{
		UserId: uint32(user.ID),
		OrderId: orderId,
		Status: statusRequest.Status,
	}

	_, err = o.Client.UpdateDriverOrderStatus(ctx, pbRequestData)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Order status updated",
		Data: pbRequestData,
	})
}