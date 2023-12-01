package repository

import "gorm.io/gorm"

type MerchantRepository struct {
	Db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return MerchantRepository{
		Db: db,
	}
}

func (m MerchantRepository) FindMultipleMenuDetails() error {
	return nil
}