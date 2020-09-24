package command

import (
	"context"
	"orderContext/domain/order"
)

type PayOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type PayOrderCommandHandler struct {
	repository order.OrderRepository
}

func NewPayOrderCommandHandler(r order.OrderRepository) PayOrderCommandHandler {
	return PayOrderCommandHandler{repository: r}
}

func (handler PayOrderCommandHandler) Handle(ctx context.Context, cmd PayOrderCommand) error {
	order := handler.repository.Get(ctx, cmd.OrderId)
	order.Pay()
	handler.repository.Update(ctx, order)
	return nil
}