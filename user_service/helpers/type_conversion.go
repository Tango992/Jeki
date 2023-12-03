package helpers

import (
	"user-service/dto"
	pb "user-service/pb/userpb"
)

func ConvertUserToUserData(user dto.UserJoinedData) *pb.UserData {
	return &pb.UserData{
		Id:        uint32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		Role:      user.Role,
		Verified:  user.Verified,
	}
}