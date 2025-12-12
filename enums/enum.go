package main

import "fmt"

// type OrderStatus intx
type OrderStatus string

const (
	// Received OrderStatus = iota
	Received OrderStatus= "received"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order Status ", status)
}
func main() {

	changeOrderStatus(Received)
}