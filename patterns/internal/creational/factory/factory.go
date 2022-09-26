// Payments method Factory, which is going to provide us with different ways of paying at a shop.
package factory

import "fmt"

const (
	Cash = iota
	DebitCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

// creation method
func GetPaymentMethod(t int) (PaymentMethod, error) {
	switch t {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf("not implemented yet")
	}
}
