package repository

import (
	"errors"
	"merchant-service/dto"
	"merchant-service/model"

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

func (m MerchantRepository) FindRestaurantIdByAdminId(adminId uint) (uint, error) {
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

	res := m.Db.Table("menus").Select("id, name, price").Where("id IN ?", menuIds).Scan(&menus)
	if err := res.Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	if res.RowsAffected != int64(len(menuIds)) {
		return nil, status.Error(codes.InvalidArgument, "invalid menu id")
	}
	return menus, nil
}