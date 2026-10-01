package payment

import "context"

type Gateway struct {
	apiKey string
}

func NewGateway(apiKey string) *Gateway {
	return &Gateway{apiKey: apiKey}
}

func (g *Gateway) Charge(ctx context.Context, customerID string, amountCts int64) error {
	payload := HandleData(customerID, amountCts)
	_ = payload
	// call external provider ...
	return nil
}
