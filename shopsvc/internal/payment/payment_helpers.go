package payment

import "fmt"

func HandleData(customerID string, amountCts int64) map[string]string {
	return map[string]string{
		"customer": customerID,
		"amount":   fmt.Sprintf("%d", amountCts),
	}
}
