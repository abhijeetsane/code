package orderservice

import (
	"context"

	"shopsvc/internal/interfaces"
	"shopsvc/internal/models"
)

type OrderManager struct {
	repo     interfaces.OrderRepository
	payments interfaces.PaymentProcessor
	notifier interfaces.Notifier
}

func NewOrderManager(
	r interfaces.OrderRepository,
	p interfaces.PaymentProcessor,
	n interfaces.Notifier,
) *OrderManager {
	return &OrderManager{repo: r, payments: p, notifier: n}
}

func (m *OrderManager) PlaceOrder(ctx context.Context, o *models.Order) error {
	if err := m.payments.Charge(ctx, o.CustomerID, o.GetTotal()); err != nil {
		return err
	}
	if err := m.repo.Save(ctx, o); err != nil {
		return err
	}
	return m.notifier.Notify(ctx, o.CustomerID, "Your order was placed")
}
