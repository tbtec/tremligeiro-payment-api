// go
package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
)

// Mock PaymentGateway
type mockPaymentGateway struct {
	createCalled bool
	createInput  *entity.Payment
	returnErr    error
}

func (m *mockPaymentGateway) Create(ctx context.Context, payment *entity.Payment) error {
	m.createCalled = true
	m.createInput = payment
	return m.returnErr
}
