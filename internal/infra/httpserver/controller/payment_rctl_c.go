package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type PaymentCreateRestController struct {
	controller *ctl.PaymentCreateController
}

func NewPaymentCreateRestController(container *container.Container) httpserver.IController {
	return &PaymentCreateRestController{
		controller: ctl.NewPaymentCreateController(container),
	}
}

func (ctl *PaymentCreateRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	paymentCheckout := dto.PaymentCheckout{}

	err := request.ParseBody(ctx, &paymentCheckout)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	err2 := ctl.controller.Execute(ctx, paymentCheckout)
	if err2 != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(dto.Payment{
		Status: dto.PaymentStatusAuthorized,
	})
}
