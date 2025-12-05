package main

import "fmt"

func forLoop() {
	var stop = 10
	for i := 1; i <= stop; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 0; i < 6; i++ {
		fmt.Print(i," ")
	}
	fmt.Println()
}


// while style

func whileStyle(){
	var count =1
	for count <5{
		count++
	}
	fmt.Println("count of ",count)
}

// do while style 

func doWhileStyle()  {
	i:=1
	for{
		i++
		if i==4{
			break
		}
	}
	fmt.Println("do while style code: ", i)
}

func main()  {
	forLoop()
	whileStyle()
	doWhileStyle()
}