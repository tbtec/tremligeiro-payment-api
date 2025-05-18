// go
package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/dto"
)

// Mock for UscPaymentWebHook
type mockUscPaymentWebHook struct {
	createCalled bool
	createInput  dto.PaymentCheckout
	createCtx    context.Context
	returnErr    error
}

// Implement the Create method
func (m *mockUscPaymentWebHook) Create(ctx context.Context, input dto.PaymentCheckout) error {
	m.createInput = input
	m.createCalled = true
	m.createCtx = ctx
	return m.returnErr
}
