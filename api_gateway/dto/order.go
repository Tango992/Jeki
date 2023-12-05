package dto

type NewOrderRequest struct {
	Address string         `json:"address" validate:"required"`
	Items   []NewOrderItem `json:"items" validate:"required"`
}

type NewOrderItem struct {
	MenuID uint32 `json:"menu_id" validate:"required"`
	Qty    uint32 `json:"qty" validate:"qty"`
}

type UpdateOrderStatus struct {
	Status string `json:"status" validate:"required,oneof=cancelled done"`
}

type XenditWebhook struct {
	ExternalId    string `json:"external_id"`
	InvoiceId     string `json:"id"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	CompletedAt   string `json:"completed_at"`
}