package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println("printing a number using variable ", num)
	var p *int = &num;
	fmt.Println("priting a number using pointer ", *p )
	fmt.Println("Printing address of variable" , p)

	var a  = &num
	fmt.Println("printing the value ", *a)

	*p = 20

	fmt.Println("update the valuse using pointer ", *p)
}