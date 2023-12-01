package service

import (
	"context"
	"order-service/model"

	xendit "github.com/xendit/xendit-go/v3"
	invoice "github.com/xendit/xendit-go/v3/invoice"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentService interface {
	MakeInvoice(primitive.ObjectID, float32) (model.Payment, error)
}

type XenditClient struct {
	Client *xendit.APIClient
}

func NewPaymentService(key string) PaymentService {
	return XenditClient{
		Client: xendit.NewClient(key),
	}
}

func (x XenditClient) MakeInvoice(externalId primitive.ObjectID, subtotal float32) (model.Payment, error) {
	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(externalId.Hex(), subtotal)
	resp, _, errXnd := x.Client.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()
	if errXnd != nil {
		return model.Payment{}, status.Error(codes.Internal, errXnd.Error())
	}
	
	paymentData := model.Payment{
		InvoiceId: *resp.Id,
		InvoiceUrl: resp.InvoiceUrl,
		Total: resp.Amount,
		Method: "pending",
		Status: resp.Status.String(),
	}
	return paymentData, nil
}