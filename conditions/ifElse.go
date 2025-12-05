package main

import "fmt"

func ifElseStatement() {
	var age = 17
	if age >= 18 {
		fmt.Println("You can vote!")
	}else{
		fmt.Println("You can't not vote!")
	}
}