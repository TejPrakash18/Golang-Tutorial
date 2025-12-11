package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tejprakash/arithmatic"
	"github.com/tejprakash/logical"
)

func main() {
	sum:=arithmatic.Add(2,3)
	fmt.Println("sum",sum)

	sub:=arithmatic.Sub(6,2)
	fmt.Println("Subtract",sub)

	greather:=logical.Greather(4,5)
	fmt.Println(greather)

	color.Red("Hello are you!")

}