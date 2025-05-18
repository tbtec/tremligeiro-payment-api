package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

// var (
// 	ErrPaymentNotFound      = xerrors.NewBusinessError("TL-PAYWH-001", "Payment not found")
// 	ErrPaymentOrderNotFound = xerrors.NewBusinessError("TL-PAYWH-002", "Order not found")
// 	ErrPaymentFinished      = xerrors.NewBusinessError("TL-PAYWH-003", "Payment already finished")
// )

type UscPaymentCreate struct {
	paymentGateway *gateway.PaymentGateway
}

func NewUseCasePaymentCreate(paymentGateway *gateway.PaymentGateway) *UscPaymentWebHook {
	return &UscPaymentWebHook{
		paymentGateway: paymentGateway,
	}
}

func (usc *UscPaymentWebHook) Create(ctx context.Context, paymentWebHook dto.PaymentCheckout) error {

	payment := entity.NewPayment(paymentWebHook.OrderId)

	err := usc.paymentGateway.Create(ctx, &payment)
	if err != nil {
		return err
	}

	return nil
}
