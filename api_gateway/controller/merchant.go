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

// Merchant      godoc
// @Summary      Get all categories
// @Description  Retrieve all restaurant datas from the database.
// @Tags         all user
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseGetAllCategories
// @Failure      400  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /categories [get]
func (m MerchantController) GetAllCategories(c echo.Context) error {
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	categories, err := m.Client.FindAllCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all categories",
		Data: categories.Categories,
	})
}

// Merchant      godoc
// @Summary      Get all restaurant datas
// @Description  Retrieve all restaurant datas from the database.
// @Tags         all user
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseGetAllRestaurant
// @Failure      400  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /restaurant [get]
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

// Merchant      godoc
// @Summary      Get restaurant by ID
// @Description  Retrieve specific restaurant data using the restaurant id.
// @Tags         all user
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseGetRestaurantByID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /restaurant/{id} [get]
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

// Merchant      godoc
// @Summary      Get menu By ID
// @Description  Retrieve specific menu data using the menu id.
// @Tags         all user
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseGetMenuById
// @Failure      400  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /menu/{id} [get]
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

// Merchant      godoc
// @Summary      Get restaurant for restaurant admin
// @Description  Retrieves restaurant data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseGetRestaurantByAdminID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/restaurant [get]
func (m MerchantController) GetRestaurantByAdminId(c echo.Context)error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	restaurantDatas, err := m.Client.FindRestaurantByAdminId(ctx, &merchantpb.AdminId{Id: uint32(user.ID)})
	if err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all restaurants",
		Data: restaurantDatas,
	})
}

// Merchant      godoc
// @Summary      Create restaurant for restaurant admin
// @Description  Creates a new restaurant data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Produce      json
// @param        request body dto.NewRestaurantData  true  "Create Restaurant"																				// Request Body
// @Success      201  {object}  dto.SwaggerResponseCreateRestaurant
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/restaurant [post]
func (m MerchantController) CreateRestaurant(c echo.Context)error{
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
	}

	var restaurantData dto.NewRestaurantData
	if err := c.Bind(&restaurantData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}
	
	if err := c.Validate(&restaurantData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	coordinate, err := m.Maps.GetCoordinate(restaurantData.Address)
	if err != nil {
		return err
	}

	pbRestaurants := &merchantpb.NewRestaurantData{
		AdminId: uint32(user.ID),
		Name: restaurantData.Name,
		Address: restaurantData.Address,
		Latitude: coordinate.Latitude,
		Longitude: coordinate.Longitude,
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

	response := dto.ResponseNewRestaurant{
		ID: uint(restaurantId.Id),
		Name: restaurantData.Name,
		Address: restaurantData.Address,
		Latitude: coordinate.Latitude,
		Longitude: coordinate.Longitude,
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Restaurant successfully created",
		Data: response,
	})
}

// Merchant      godoc
// @Summary      Update restaurant for restaurant admin
// @Description  Updates existing restaurant data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Produce      json
// @param        request body dto.UpdateRestaurantData  true  "Update Restaurant"
// @Success      200  {object}  dto.SwaggerResponseUpdateRestaurant
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/restaurant [put]
func (m MerchantController) UpdateRestaurant(c echo.Context) error{
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
	}

	var restaurantUpdate dto.UpdateRestaurantData
	if err := c.Bind(&restaurantUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&restaurantUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	coordinate, err := m.Maps.GetCoordinate(restaurantUpdate.Address)
	if err != nil {
		return err
	}

	pbUpdateRestaurant :=  &merchantpb.UpdateRestaurantData{
		AdminId: uint32(user.ID),
		Name: restaurantUpdate.Name,
		Address: restaurantUpdate.Address,
		Latitude: coordinate.Latitude,
		Longitude: coordinate.Longitude,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()
	
	if _, err := m.Client.UpdateRestaurant(ctx, pbUpdateRestaurant); err != nil {
		return helpers.AssertGrpcStatus(err)
	}

	response := dto.ResponseUpdateRestaurant{
		Name: restaurantUpdate.Name,
		Address: restaurantUpdate.Address,
		Latitude: coordinate.Latitude,
		Longitude: coordinate.Longitude,
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Restaurant successfully updated",
		Data: response,
	})
}

// Merchant      godoc
// @Summary      Get menu for restaurant admin
// @Description  Retrieves restaurant menus specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseGetMenuByAdminID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/menu [get]
func (m MerchantController) GetMenuByAdminId(c echo.Context)error{
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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
		Message: "Get all menus",
		Data: menu,
	})
}

// Merchant      godoc
// @Summary      Get one menu for restaurant admin
// @Description  Retrieves one menu specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseGetMenuIdByAdminID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/menu/{id} [get]
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
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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

// Merchant      godoc
// @Summary      Create menu for restaurant admin
// @Description  Creates new menu data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Accept       json
// @Produce      json
// @param        request body dto.SwaggerRequestMenu  true  "Create Menu"
// @Success      201  {object}  dto.SwaggerResponseCreateMenuByAdminID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/menu [post]
func (m MerchantController) CreateMenu(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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

// Merchant      godoc
// @Summary      Update menu for restaurant admin
// @Description  Updates existing menu data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Accept       json
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @param 		request body dto.SwaggerRequestMenu  true  "Update Menu"
// @Success      200  {object}  dto.SwaggerResponseUpdateMenuByAdminID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/menu/{id} [put]
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
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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


// Merchant      godoc
// @Summary      Delete menu for restaurant admin
// @Description  Deletes existing menu for the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         merchant
// @Accept       json
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /merchant/menu/{id} [delete]
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

	if user.Role != adminRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Only admin role is allowed"))
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