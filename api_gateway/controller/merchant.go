package controller

import (
	"api-gateway/dto"
	"api-gateway/helpers"
	"api-gateway/pb/merchantpb"
	"api-gateway/service"
	"api-gateway/utils"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/labstack/echo/v4"
)

type MerchantController struct{
	Client merchantpb.MerchantClient
	Maps service.Maps
}

func NewMerchantController(client merchantpb.MerchantClient, ms service.Maps) MerchantController {
	return MerchantController{
		Client: client,
		Maps: ms,
	}
}

func (m MerchantController) TestMap(c echo.Context) error {
	var request struct{Address string `json:"address" validate:"required"`}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	result, err := m.Maps.GetCoordinate(request.Address)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get Coordinate",
		Data: result,
	})
}

func (m MerchantController) GetAllRestaurantsForCustomer(c echo.Context) error {
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	restaurantDatas, err := m.Client.FindAllRestaurants(ctx, &emptypb.Empty{})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all restaurants",
		Data: restaurantDatas,
	})
}

func (m MerchantController) GetRestaurantById(c echo.Context) error {
	restaurantIdTmp := c.Param("id")
	restaurantId, err := strconv.Atoi(restaurantIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	pbRestaurantId := &merchantpb.IdRestaurant{
		Id: uint32(restaurantId),
	}
	
	restaurantData, err := m.Client.FindRestaurantById(ctx, pbRestaurantId)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}
	
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get restaurant by ID",
		Data: restaurantData,
	})
}

func (m MerchantController) GetMenuById(c echo.Context) error {
	menuIdTmp := c.Param("id")
	menuId, err := strconv.Atoi(menuIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	pbMenuId := &merchantpb.MenuId{
		Id: uint32(menuId),
	}

	menu, err := m.Client.FindMenuById(ctx, pbMenuId)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}
	
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get menu by ID",
		Data: menu,
	})
}

func (m MerchantController) GetAllRestaurants(c echo.Context)error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	restaurantDatas, err := m.Client.FindAllRestaurants(ctx, &emptypb.Empty{})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all restaurants",
		Data: restaurantDatas,
	})
}

func (m MerchantController) CreateRestaurant(c echo.Context)error{
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}

	var restaurantData dto.NewRestaurantData
	if err := c.Bind(&restaurantData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	if err := c.Validate(&restaurantData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbRestaurants := &merchantpb.NewRestaurantData{
		AdminId: uint32(user.ID),
		Name: restaurantData.Name,
		Address: restaurantData.Address,
		Latitude: 0.1,						// Temporary
		Longitude: 0.1,						// Temporary
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	restaurantId, err := m.Client.CreateRestaurant(ctx, pbRestaurants)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Restaurant successfully posted",
		Data: restaurantId,
	})
}

func (m MerchantController) UpdateRestaurant(c echo.Context) error{
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}

	var restauranUpdate dto.UpdateRestaurantData
	if err := c.Bind(&restauranUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&restauranUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbUpdateRestauran :=  &merchantpb.UpdateRestaurantData{
		AdminId: uint32(user.ID),
		Name: restauranUpdate.Name,
		Address: restauranUpdate.Address,
		Latitude: 0.1,						// Temporary
		Longitude: 0.1,						// Temporary
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()
	
	if _, err := m.Client.UpdateRestaurant(ctx, pbUpdateRestauran); err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Restauran successfully updated",
		Data: restauranUpdate,
	})


}

func (m MerchantController) GetMenu(c echo.Context)error{
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	menu, err := m.Client.FindMenusByAdminId(ctx, &merchantpb.AdminId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}
	
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all menu ",
		Data: menu,
	})

}

func (m MerchantController) GetOneMenuByAdminId(c echo.Context) error {
	menuIdTmp := c.Param("id")
	menuId, err := strconv.Atoi(menuIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	pbRequestData := &merchantpb.AdminIdMenuId{
		AdminId: uint32(user.ID),
		MenuId: uint32(menuId),
	}
	
	menu, err := m.Client.FindOneMenuByAdminId(ctx, pbRequestData)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}
	
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get menu by admin ID",
		Data: menu,
	})
}

func (m MerchantController) CreateMenu(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}
	
	var menuData dto.NewMenuData
	if err := c.Bind(&menuData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	if err := c.Validate(&menuData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbMenuData := &merchantpb.NewMenuData{
		AdminId: uint32(user.ID),
		Name: menuData.Name,
		CategoryId: menuData.CategoryId,
		Price: menuData.Price,
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()
	
	menuId, err := m.Client.CreateMenu(ctx, pbMenuData)
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}
	menuData.ID = menuId.Id
	
	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Menu successfully posted",
		Data: menuData,
	})
}

func (m MerchantController) UpdateMenu(c echo.Context) error {
	menuIdTmp := c.Param("id")
	menuId, err := strconv.Atoi(menuIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Credential has to be admin"))
	}

	var menuData dto.UpdateMenuData
	if err := c.Bind(&menuData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	if err := c.Validate(&menuData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	menuData.MenuId = uint32(menuId)

	pbUpdateMenu := &merchantpb.UpdateMenuData{
		AdminId: uint32(user.ID),
		MenuId: menuData.MenuId,
		Name: menuData.Name,
		CategoryId: menuData.CategoryId,
		Price: menuData.Price,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()
	
	if _, err := m.Client.UpdateMenu(ctx, pbUpdateMenu); err != nil {
		return helpers.AssertGrpcStatus(err)
	}
	
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Menu successfully updated",
		Data: menuData,
	})
}

func (m MerchantController) DeleteMenu(c echo.Context) error {
	menuIdTmp := c.Param("id")
	menuId, err := strconv.Atoi(menuIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}
	
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()
	
	pbRequestData := &merchantpb.AdminIdMenuId{
		AdminId: uint32(user.ID),
		MenuId: uint32(menuId),
	}
	
	if _, err := m.Client.DeleteMenu(ctx, pbRequestData); err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Menu successfully deleted",
		Data: fmt.Sprintf("Deleted menu on ID = %v", menuId),
	})
}
