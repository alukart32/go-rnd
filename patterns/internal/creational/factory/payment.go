package factory

import "fmt"

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (o *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (o *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

func (o *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (credit card)\n", amount)
}
