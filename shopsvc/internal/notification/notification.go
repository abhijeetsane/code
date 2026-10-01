package notification

import "context"

type EmailSender struct{}

func (e *EmailSender) Notify(ctx context.Context, customerID, msg string) error {
	return SendNotification(customerID, msg)
}

func SendNotification(customerID, msg string) error {
	// send email ...
	return nil
}
