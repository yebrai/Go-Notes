package behavioral

import "fmt"

// Strategy interface
type PaymentStrategy interface {
	Pay(amount int)
}

// Concrete strategies
type CreditCard struct{}

func (cc *CreditCard) Pay(amount int) {
	fmt.Println("Paid", amount, "using credit card")
}

type PayPal struct{}

func (pp *PayPal) Pay(amount int) {
	fmt.Println("Paid", amount, "using PayPal")
}

// Context
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	pp.strategy = strategy
}

func (pp *PaymentProcessor) Pay(amount int) {
	pp.strategy.Pay(amount)
}

func main() {
	processor := &PaymentProcessor{}

	processor.SetStrategy(&CreditCard{})
	processor.Pay(100)

	processor.SetStrategy(&PayPal{})
	processor.Pay(200)
}
