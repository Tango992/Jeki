package controller

import "api-gateway/pb/merchantpb"

type MerchantController struct{
	Client merchantpb.MerchantClient
}

func NewMerchantController(client merchantpb.MerchantClient) MerchantController {
	return MerchantController{
		Client: client,
	}
}