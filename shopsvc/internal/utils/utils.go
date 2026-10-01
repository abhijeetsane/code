package utils

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidInput = errors.New("invalid input")

func ValidateEmail(s string) bool {
	return strings.Contains(s, "@")
}

func FormatMoney(cents int64) string {
	return fmt.Sprintf("₹%d.%02d", cents/100, cents%100)
}

func Retry(n int, fn func() error) error {
	var err error
	for i := 0; i < n; i++ {
		if err = fn(); err == nil {
			return nil
		}
	}
	return err
}
