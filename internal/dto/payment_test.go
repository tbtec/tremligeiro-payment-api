package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentCheckoutMarshalUnmarshal(t *testing.T) {
	meta := MetaData{
		PaymentId:         strPtr("pid123"),
		PaymentWebHookUrl: strPtr("https://webhook.url"),
	}
	checkout := PaymentCheckout{
		OrderId:     "order-1",
		TotalAmount: 99.99,
		Products: []PaymentItemCheckoutProduct{
			{ProductId: "prod-1", Quantity: 2},
			{ProductId: "prod-2", Quantity: 1},
		},
		MetaData: meta,
	}

	data, err := json.Marshal(checkout)
	assert.NoError(t, err)

	var unmarshaled PaymentCheckout
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, checkout.OrderId, unmarshaled.OrderId)
	assert.Equal(t, checkout.TotalAmount, unmarshaled.TotalAmount)
	assert.Len(t, unmarshaled.Products, 2)
	assert.Equal(t, "prod-1", unmarshaled.Products[0].ProductId)
	assert.Equal(t, 2, unmarshaled.Products[0].Quantity)
	assert.NotNil(t, unmarshaled.MetaData.PaymentId)
	assert.Equal(t, "pid123", *unmarshaled.MetaData.PaymentId)
	assert.NotNil(t, unmarshaled.MetaData.PaymentWebHookUrl)
	assert.Equal(t, "https://webhook.url", *unmarshaled.MetaData.PaymentWebHookUrl)
}

func TestPaymentItemCheckoutProductJSON(t *testing.T) {
	item := PaymentItemCheckoutProduct{
		ProductId: "prod-xyz",
		Quantity:  5,
	}
	data, err := json.Marshal(item)
	assert.NoError(t, err)

	var unmarshaled PaymentItemCheckoutProduct
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, item.ProductId, unmarshaled.ProductId)
	assert.Equal(t, item.Quantity, unmarshaled.Quantity)
}

func strPtr(s string) *string {
	return &s
}
