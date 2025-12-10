package main

import (
	"fmt"
)

type paymenter interface{
	pay(amount float32)
}
type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
//  rozarpayPatment:= rozarpay{}
//  rozarpayPatment.pay(100)

// stripPayment := strip{}
// stripPayment.pay(100)
	p.gateway.pay(amount)
}

type rozarpay struct{}

func (r rozarpay) pay(amount float32) {
	fmt.Println("Making payment using rozarpay, ", amount)
}

type strip struct{}

func (s strip)pay(amount float32){
	fmt.Println("Making payment using strip ", amount)
}
func main() {
	// newPayment := payment{}
	// newPayment.makePayment(100.00)

	// payByStrip := strip{}
	// payByStrip.pay(100.10)
	payByRozar := rozarpay{}
	// payByRozar.pay(100)


	newPayment := payment{
		gateway:  payByRozar,
	}
	newPayment.makePayment(100)
}