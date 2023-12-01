package controller

import (
	"context"
	pb "merchant-service/pb/merchantpb"
	"merchant-service/repository"
)

type Server struct {
	pb.UnimplementedMerchantServer
	Repository repository.Merchant
}

func NewUserController(r repository.Merchant) Server {
	return Server{
		Repository: r,
	}
}

func (s Server) FindMenuDetailsWithSubtotal(ctx context.Context, data *pb.RequestMenuDetails) (*pb.ResponseMenuDetails, error) {
	return nil, nil
}