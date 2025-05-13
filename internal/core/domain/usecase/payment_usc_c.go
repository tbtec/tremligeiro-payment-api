package usecase

import (
	"context"

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

	// payment, err := usc.paymentGateway.FindOne(ctx, paymentWebHook.PaymentId)
	// if err != nil {
	// 	return ErrPaymentNotFound
	// }
	// if payment.IsFinished() {
	// 	return ErrPaymentFinished
	// }

	// order, err := usc.orderGateway.FindOne(ctx, payment.OrderId)
	// if err != nil {
	// 	return ErrPaymentOrderNotFound
	// }

	// if paymentWebHook.Status == "approved" {
	// 	order.SetStatus(entity.OrderStatusReceived)
	// 	payment.SetStatus(entity.PaymentStatusAuthorized)
	// 	usc.orderGateway.Update(ctx, order)
	// 	usc.paymentGateway.Update(ctx, payment)
	// }

	// if paymentWebHook.Status == "repproved" {
	// 	payment.SetStatus(entity.PaymentStatusNotAuthorized)
	// 	usc.paymentGateway.Update(ctx, payment)
	// }

	return nil
}
