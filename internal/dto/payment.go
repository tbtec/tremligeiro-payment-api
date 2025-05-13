package dto

type PaymentWebHook struct {
	Resource  string `json:"resource"`
	Topic     string `json:"topic"`
	Status    string `json:"status"`
	PaymentId string `json:"paymentId"`
}

type Payment struct {
	ID          string   `json:"id"`
	CustomerId  *string  `json:"customerId"`
	TotalAmount float64  `json:"totalAmount"`
	MetaData    MetaData `json:"metadata"`
}

type PaymentCheckout struct {
	OrderId  string                       `json:"orderId" validate:"required"`
	Products []PaymentItemCheckoutProduct `json:"products" validate:"required"`
	MetaData MetaData                     `json:"metadata"`
}

type PaymentItemCheckoutProduct struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
