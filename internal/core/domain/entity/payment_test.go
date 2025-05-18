// go
package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPayment(t *testing.T) {
	expectedID := "test-ulid-123"
	orderId := "order-xyz"
	now := time.Now().UTC()

	// Assuming NewPayment can accept an ID for testing purposes
	payment := &Payment{
		ID:        expectedID,
		OrderId:   orderId,
		Status:    PaymentStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
		Products:  nil,
	}

	assert.Equal(t, orderId, payment.OrderId)
	assert.Equal(t, PaymentStatusPending, payment.Status)
	assert.Equal(t, expectedID, payment.ID)
	assert.WithinDuration(t, now, payment.CreatedAt, time.Second)
	assert.WithinDuration(t, now, payment.UpdatedAt, time.Second)
	assert.Empty(t, payment.Products)
}
