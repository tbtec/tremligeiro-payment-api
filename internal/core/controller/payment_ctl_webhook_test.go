// go
package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

type mockPaymentService struct{}
type mockPaymentRepository struct{}

func (m *mockPaymentRepository) Create(ctx context.Context, payment *model.Payment) error {
	return nil
}

// Add the missing FindByOrderId method to satisfy repository.IPaymentRepository
func (m *mockPaymentRepository) FindByOrderId(ctx context.Context, orderId string) (*model.Payment, error) {
	return nil, nil
}

// Add the missing FindOne method to satisfy repository.IPaymentRepository
func (m *mockPaymentRepository) FindOne(ctx context.Context, id string) (*model.Payment, error) {
	return nil, nil
}

// Add the missing Update method to satisfy repository.IPaymentRepository
func (m *mockPaymentRepository) Update(ctx context.Context, payment *model.Payment) error {
	return nil
}

// Implement the correct RequestPayment method to satisfy external.IPaymentService
func (m *mockPaymentService) RequestPayment(ctx context.Context, req external.PaymentRequest) (external.PaymentResponse, error) {
	return external.PaymentResponse{}, nil
}

func TestNewPaymentWebHookController(t *testing.T) {
	mockContainer := &container.Container{
		PaymentService:    &mockPaymentService{},
		PaymentRepository: &mockPaymentRepository{},
	}

	controller := NewPaymentWebHookController(mockContainer)

	assert.NotNil(t, controller)
	assert.NotNil(t, controller.usc)
}
