package models

import "time"

type Order struct {
	ID         string
	CustomerID string
	Items      []LineItem
	CreatedAt  time.Time
}

type LineItem struct {
	SKU      string
	Quantity int
	PriceCts int64
}

type Customer struct {
	ID    string
	Email string
	Name  string
}

func (o *Order) GetTotal() int64 {
	var total int64
	for _, it := range o.Items {
		total += it.PriceCts * int64(it.Quantity)
	}
	return total
}
