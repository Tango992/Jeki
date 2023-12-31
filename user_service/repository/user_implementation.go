package repository

import (
	"errors"
	"user-service/dto"
	"user-service/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	onlineDriverStatusID  = 1
	ongoingDriverStatusID = 2
	offlineDriverStatusID = 3
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		Db: db,
	}
}

func (u UserRepository) GetUserData(email string) (dto.UserJoinedData, error) {
	var userData dto.UserJoinedData

	result := u.Db.Table("users u").
		Select("u.id, u.first_name, u.last_name, u.email, u.password, u.birth_date, u.created_at, r.name AS role, v.validate AS verified").
		Where("u.email = ?", email).
		Joins("JOIN roles r ON u.role_id = r.id").
		Joins("JOIN verifications v ON v.user_id = u.id").
		Take(&userData)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserJoinedData{}, status.Error(codes.NotFound, err.Error())
		}
		return userData, status.Error(codes.Internal, err.Error())
	}
	return userData, nil
}

func (u UserRepository) CreateUser(data *models.User) error {
	result := u.Db.Create(data)
	if err := result.Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return status.Error(codes.AlreadyExists, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (u UserRepository) AddToken(data *models.Verification) error {
	if err := u.Db.Create(data).Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (u UserRepository) GetAvailableDriver() (dto.DriverData, error) {
	var driverData dto.DriverData
	result := u.Db.Table("drivers d").
		Select("d.user_id AS id, u.first_name || ' ' || u.last_name AS name").
		Joins("JOIN users u on d.user_id = u.id").
		Joins("JOIN driver_statuses ds on d.driver_status_id = ds.id").
		Where("ds.status = ?", "available").
		Take(&driverData)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.DriverData{}, status.Error(codes.Unavailable, "no available driver to pick up the order")
		}
		return dto.DriverData{}, status.Error(codes.Internal, err.Error())
	}

	if err := u.SetDriverStatusOngoing(uint(driverData.ID)); err != nil {
		return dto.DriverData{}, err
	}
	return driverData, nil
}

func (u UserRepository) SetDriverStatusOnline(driverID uint) error {
	result := u.Db.Table("drivers").
		Where("user_id = ?", driverID).
		Update("driver_status_id", onlineDriverStatusID)

	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.InvalidArgument, "Invalid driver ID")
	}
	return nil
}

func (u UserRepository) SetDriverStatusOngoing(driverID uint) error {
	result := u.Db.Table("drivers").
		Where("user_id = ?", driverID).
		Update("driver_status_id", ongoingDriverStatusID)

	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.InvalidArgument, "Invalid driver ID")
	}
	return nil
}

func (u UserRepository) SetDriverStatusOffline(driverID uint) error {
	result := u.Db.Table("drivers").
		Where("user_id = ?", driverID).
		Update("driver_status_id", offlineDriverStatusID)

	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.InvalidArgument, "Invalid driver ID")
	}
	return nil
}

func (u UserRepository) VerifyNewUser(userID uint32, token string) error {
	verificationData := models.Verification{
		UserID: uint(userID),
	}

	res := u.Db.Model(&verificationData).
		Where("token = ?", token).
		Update("validate", true)
	
	if err := res.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if res.RowsAffected == 0 {
		return status.Error(codes.Unauthenticated, "invalid verification credentials")
	}
	return nil
}

func (u UserRepository) CreateDriverData(userID uint32) error {
	driverData := models.Driver{
		UserID: uint(userID),
		DriverStatusID: offlineDriverStatusID,
	}
	
	res := u.Db.Create(&driverData)
	if err := res.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}