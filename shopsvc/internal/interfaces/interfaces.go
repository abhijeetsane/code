package interfaces

import (
	"context"

	"shopsvc/internal/models"
)

type OrderRepository interface {
	Save(ctx context.Context, o *models.Order) error
	FindByID(ctx context.Context, id string) (*models.Order, error)
	ListByCustomer(ctx context.Context, customerID string) ([]*models.Order, error)
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id, status string) error
}

type PaymentProcessor interface {
	Charge(ctx context.Context, customerID string, amountCts int64) error
}

type Notifier interface {
	Notify(ctx context.Context, customerID, msg string) error
}
