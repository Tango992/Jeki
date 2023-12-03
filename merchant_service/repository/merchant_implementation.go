package repository

import (
	"errors"
	"merchant-service/dto"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type MerchantRepository struct {
	Db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return MerchantRepository{
		Db: db,
	}
}

func (m MerchantRepository) FindAllRestaurants() ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	if err := m.Db.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (m MerchantRepository) FindRestaurantByID(id uint32) (dto.RestaurantDataCompact, error) {
    var restaurant dto.RestaurantDataCompact
    if err := m.Db.Table("restaurants").Where("id = ?", id).Take(&restaurant).Error; err != nil {
        return dto.RestaurantDataCompact{}, err
    }
    return restaurant, nil
}

func (m MerchantRepository) CreateRestaurant(data *model.Restaurant) error {
	if err := m.Db.Create(data).Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "idx_restaurants_admin_id" (SQLSTATE 23505)` {
			return status.Error(codes.FailedPrecondition, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (m MerchantRepository) CreateMenu(data *model.Menu) error {
	if err := m.Db.Create(data).Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (m MerchantRepository) FindRestaurantIdByAdminId(adminId uint32) (uint, error) {
	var restaurantId uint
	if err := m.Db.Table("restaurants").Select("id").Where("admin_id = ?", adminId).Take(&restaurantId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, status.Error(codes.NotFound, err.Error())
		}
		return 0, status.Error(codes.Internal, err.Error())
	}
	return restaurantId, nil
}

func (m MerchantRepository) DeleteMenu(restoId, menuId uint) error {
	restoData := model.Menu{
		ID: menuId,
	}

	res := m.Db.Delete(&restoData, "restaurant_id = ?", restoId)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "invalid menu id")
	}
	return nil
}

func (m MerchantRepository) FindMultipleMenuDetails(menuIds []int) ([]dto.MenuTmp, error) {
	var menus []dto.MenuTmp

	res := m.Db.Table("menus").Select("id, name, price").Where("id IN ?", menuIds).Order("id").Scan(&menus)
	if err := res.Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	if res.RowsAffected != int64(len(menuIds)) {
		return nil, status.Error(codes.InvalidArgument, "invalid menu id")
	}
	return menus, nil
}

func (m MerchantRepository) FindMenusByAdminId(adminId uint32) ([]*pb.MenuCompact, error) {
	menus := []*pb.MenuCompact{}

	res := m.Db.Raw(`
		SELECT m.id, m.name, c.name AS category, m.price
		FROM menus m
		JOIN categories c ON c.id = m.category_id
		WHERE m.restaurant_id = 
			(SELECT id
			FROM restaurants r
			WHERE r.admin_id = ?
			LIMIT 1)`, adminId).
		Scan(&menus)

	if err := res.Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return menus, nil
}

func (m MerchantRepository) FindOneMenuByAdminId(menuId, adminId uint32) (*pb.MenuCompact, error) {
	var menu *pb.MenuCompact

	res := m.Db.Raw(`
	SELECT m.name, c.name AS category, m.price
	FROM menus m
	JOIN categories c ON c.id = m.category_id
	WHERE m.id = ? AND m.restaurant_id = 
		(SELECT id
		FROM restaurants r
		WHERE r.admin_id = ?
		LIMIT 1)`, menuId, adminId).
	Take(&menu)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return menu, nil
}

func (m MerchantRepository) FindRestaurantByAdminId(adminId uint32) (*pb.RestaurantData, error) {
	var restaurant *pb.RestaurantData

	res := m.Db.Table("restaurants").
		Where("admin_id = ?", adminId).
		Take(&restaurant)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return restaurant, nil
}

func (m MerchantRepository) FindMenuByRestaurantId(restaurantId uint32) ([]*pb.Menu, error) {
	menus := []*pb.Menu{}

	res := m.Db.Table("menus m").
		Select("m.id, m.name, c.name AS category, m.price").
		Joins("JOIN categories c ON c.id = m.category_id").
		Where("m.restaurant_id = ?", restaurantId).
		Scan(&menus)

	if err := res.Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return menus, nil
}

func (m MerchantRepository) FindMenuById(menuId uint32) (*pb.Menu, error) {
	var menu *pb.Menu

	res := m.Db.Table("menus m").
		Select("m.id, m.name, c.name AS category, m.price").
		Joins("JOIN categories c ON c.id = m.category_id").
		Where("m.id = ?", menuId).
		Take(&menu)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return menu, nil
}

func (m MerchantRepository) UpdateMenu(data *pb.UpdateMenuData) error {
	menu := model.Menu{ID: uint(data.MenuId)}
	
	res := m.Db.Model(&menu).
		Updates(model.Menu{
			Name: data.Name,
			CategoryId: uint(data.CategoryId),
			Price: data.Price,
		})

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (m MerchantRepository) UpdateRestaurant(restaurantId uint, data *pb.UpdateRestaurantData) error {
	restaurant := model.Restaurant{ID: uint(restaurantId)}

	res := m.Db.Model(&restaurant).
		Updates(model.Restaurant{
			Name: data.Name,
			Address: data.Address,
			Latitude: data.Latitude,
			Longitude: data.Longitude,
		})
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (m MerchantRepository) FindAdminIdByMenuId(menuId uint32) (uint32, error) {
	var adminId uint32
	res := m.Db.Table("menus m").
		Select("r.admin_id").
		Joins("JOIN restaurants r ON r.id = m.restaurant_id").
		Where("m.id = ?", menuId).
		Take(&adminId)
	
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, status.Error(codes.NotFound, err.Error())
		}
		return 0, status.Error(codes.Internal, err.Error())
	}
	return adminId, nil
}

func (m MerchantRepository) FindRestaurantMetadataByMenuIds(menuIds []int) (*pb.RestaurantMetadata, error) {
	var restaurantMetadata []*pb.RestaurantMetadata

	res := m.Db.Raw(`
		SELECT r.id, r.admin_id, r.name, r.latitude, r.longitude
		FROM restaurants r
		JOIN menus m ON m.restaurant_id = r.id
		WHERE m.id IN ?
		GROUP BY r.id`, menuIds).
		Scan(&restaurantMetadata)
	
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	if res.RowsAffected != 1 {
		return nil, status.Error(codes.InvalidArgument, "menus can only be from one restaurant per order")
	}
	return restaurantMetadata[0], nil
}