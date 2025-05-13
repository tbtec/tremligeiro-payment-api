package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type PaymentCreateController struct {
	usc *usecase.UscPaymentWebHook
}

func NewPaymentCreateController(container *container.Container) *PaymentCreateController {
	return &PaymentCreateController{
		usc: usecase.NewUseCasePaymentCreate(
			gateway.NewPaymentGateway(container.PaymentService, container.PaymentRepository)),
	}
}

func (ctl *PaymentCreateController) Execute(ctx context.Context, input dto.PaymentCheckout) error {
	return ctl.usc.Create(ctx, input)
}
