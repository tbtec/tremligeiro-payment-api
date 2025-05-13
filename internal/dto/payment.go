package dto

type PaymentWebHook struct {
	Resource  string `json:"resource"`
	Topic     string `json:"topic"`
	Status    string `json:"status"`
	PaymentId string `json:"paymentId"`
}
type PaymentCheckout struct {
	OrderId     string                       `json:"orderId" validate:"required"`
	TotalAmount float64                      `json:"totalAmount"`
	Products    []PaymentItemCheckoutProduct `json:"products" validate:"required"`
	MetaData    MetaData                     `json:"metadata"`
}

type PaymentItemCheckoutProduct struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type Payment struct {
	Status PaymentStatus `json:"status"`
}

type PaymentStatus string

const (
	PaymentStatusAuthorized    PaymentStatus = "AUTHORIZED"
	PaymentStatusNotAuthorized PaymentStatus = "NOT_AUTHORIZED"
)
