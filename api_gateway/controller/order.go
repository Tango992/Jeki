package controller

import (
	"api-gateway/dto"
	"api-gateway/helpers"
	"api-gateway/pb/orderpb"
	"api-gateway/service"
	"api-gateway/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Client orderpb.OrderServiceClient
	Maps   service.Maps
}

func NewOrderController(client orderpb.OrderServiceClient, maps service.Maps) OrderController {
	return OrderController{
		Client: client,
		Maps:   maps,
	}
}

// @Summary 	Create a new order by user
// @Description Create a new order for the logged-in user. You will need an 'Authorization' cookie attached with this request.
// @Tags 		customer
// @Accept 		json
// @Produce 	json
// @Param 		orderRequest body dto.NewOrderRequest true "Order details"
// @Success 	201 {object} dto.SwaggerResponseOrder
// @Failure 	400 {object} utils.ErrResponse
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/orders [post]
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
		Name:   user.Name,
		Email:  user.Email,
		Address: &orderpb.Address{
			Latitude:  coordinate.Latitude,
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
		Data:    response,
	})
}

// @Summary 	Get all orders for a user
// @Description Get all orders for the logged-in user. You will need an 'Authorization' cookie attached with this request.
// @Tags 		customer
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.SwaggerResponese
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/orders [get]
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
		Data:    orders,
	})
}

// @Summary 	Get ongoing orders for a user
// @Description Get ongoing orders for the logged-in user. You will need an 'Authorization' cookie attached with this request.
// @Tags 		customer
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.SwaggerResponese
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/ongoing [get]
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
		Data:    orders,
	})
}

// Orders       godoc
// @Summary 	Get order by ID
// @Description Get order details by order ID. You will need an 'Authorization' cookie attached with this request.
// @Tags 		customer
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Order ID"
// @Success 	200 {object} dto.SwaggerResponseOrder
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/orders/{id} [get]
func (o OrderController) UsersGetOrderById(c echo.Context) error {
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
		Data:    order,
	})
}

// @Summary      Get order by ID
// @Description Get order details by order ID. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Accept       json
// @Produce      json
// @Param 		 id   path      string  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseOrder
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/orders/{id} [get]
func (o OrderController) MerchantGetOrderById(c echo.Context) error {
	orderId := c.Param("id")

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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
		Data:    order,
	})
}

// @Summary      Get order by ID
// @Description  You will need an 'Authorization' cookie attached with this request.
// @Tags         driver
// @Accept       json
// @Produce      json
// @Param 		 id   path      string  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseOrder
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /driver/orders/{id} [get]
func (o OrderController) DriverGetOrderById(c echo.Context) error {
	orderId := c.Param("id")

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

	order, err := o.Client.GetOrderById(ctx, &orderpb.OrderId{Id: orderId})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get order by ID",
		Data:    order,
	})
}

// @summary 	Get all orders for a merchant
// @description Get all orders for the logged-in merchant (admin). You will need an 'Authorization' cookie attached with this request.
// @Tags 		merchant
// @Produce 	json
// @Success 	200 {object} dto.SwaggerResponese
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/merchant/orders [get]
func (o OrderController) MerchantGetAllOrders(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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
		Data:    orders,
	})
}

// @Summary 	Get ongoing orders for a merchant
// @Description Get ongoing orders for the logged-in merchant (admin). You will need an 'Authorization' cookie attached with this request.
// @Tags 		merchant
// @Produce 	json
// @Success 	200 {object} dto.SwaggerResponese
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/merchant/ongoing [get]
func (o OrderController) MerchantGetOngoingOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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
		Data:    orders,
	})
}

// Orders        godoc
// @Summary      Update Order from Merchants
// @Description  You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Accept       json
// @Produce      json
// @Param 		 id   path      string  true  "Id"
// @param 		request body dto.UpdateOrderStatus  true  "Merchant Update Order"
// @Success      200  {object}  dto.SwaggerResponseUpdateOrder
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/orders/{id} [put]
func (o OrderController) MerchantUpdateOrder(c echo.Context) error {
	orderId := c.Param("id")

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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
		UserId:  uint32(user.ID),
		OrderId: orderId,
		Status:  statusRequest.Status,
	}

	_, err = o.Client.UpdateRestaurantOrderStatus(ctx, pbRequestData)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Order status updated",
		Data:    pbRequestData,
	})
}

// Orders        godoc
// @Summary      Driver Get All Orders
// @Description  You will need an 'Authorization' cookie attached with this request.
// @Tags         driver
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseDriverGetAllOrders
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /driver/orders [get]
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
		Data:    orders,
	})
}

// Orders        godoc
// @Summary      Driver Get Current Order
// @Description  You will need an 'Authorization' cookie attached with this request.
// @Tags         driver
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseDriverGetCurrentOrder
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /driver/ongoing [get]
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
		Data:    order,
	})
}

// Orders        godoc
// @Summary      Driver Update order
// @Description  You will need an 'Authorization' cookie attached with this request.
// @Tags         driver
// @Accept       json
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @param 		request body dto.UpdateOrderStatus  true  "Driver Update Order"
// @Success      200  {object}  dto.SwaggerResponseUpdateOrder
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /driver/orders/{id} [put]
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
		UserId:  uint32(user.ID),
		OrderId: orderId,
		Status:  statusRequest.Status,
	}

	_, err = o.Client.UpdateDriverOrderStatus(ctx, pbRequestData)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Order status updated",
		Data:    pbRequestData,
	})
}

func (o OrderController) PaymentUpdate(c echo.Context) error {
	webhookToken := c.Request().Header.Get("x-callback-token")
	if webhookToken != os.Getenv("XENDIT_WEBHOOK_TOKEN") {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid webhook token"))
	}

	var paymentData dto.XenditWebhook
	if err := c.Bind(&paymentData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbRequestData := &orderpb.RequestUpdatePayment{
		OrderId:     paymentData.ExternalId,
		InvoiceId:   paymentData.InvoiceId,
		Method:      paymentData.PaymentMethod,
		Status:      paymentData.Status,
		CompletedAt: paymentData.CompletedAt,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	if _, err := o.Client.UpdatePaymentOrderStatus(ctx, pbRequestData); err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.NoContent(http.StatusOK)
}
