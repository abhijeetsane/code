package customers

import (
	"context"

	"shopsvc/internal/models"
	"shopsvc/internal/utils"
)

type CustomerService struct {
	// ...
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) Register(ctx context.Context, c *models.Customer) error {
	if !utils.ValidateEmail(c.Email) {
		return utils.ErrInvalidInput
	}
	// ...
	return nil
}
