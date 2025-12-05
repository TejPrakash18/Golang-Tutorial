package main

import "fmt"

func elseIfStatement() {
	var number1 = 70

	if number1 == 70 {
		fmt.Println("you will be Passed!")
	}else if number1 >= 80 && number1 < 90{
		fmt.Println("you will be passed with A Grade")
	}else{
		fmt.Println("You will be passed with A+ Grade")
	}
}