package main

import "fmt"

type Paymet interface {
	Pay()
}

type CashPyment struct{}

func (CashPyment) Pay() {
	fmt.Printf("Pyment using cash")
}

func ProcessPayment(p Paymet) {
	p.Pay()
}

type BackPayment struct{}

type BackPaymentAdapter struct {
	BackPayment *BackPayment
	bankAccount int
}

func (bpa *BackPaymentAdapter) pay() {
	bpa.BackPayment.Pay(bpa.bankAccount)
}

func (BackPayment) Pay(banlCount int) {
	fmt.Printf("bankAccount %v", banlCount)
}

func main() {
	cash := &CashPyment{}
	ProcessPayment(cash)

	bpa := &BackPaymentAdapter{
		bankAccount: 10,
		BackPayment: &BackPayment{},
	}
	ProcessPayment(bpa)
}
