package main

import (
	"fmt"
	"time"
)

// 2nd goroutine 
func processNum(numChan chan int){
	fmt.Println(<-numChan)
}


// send the data from one goroutine to another 

func sum(result chan int, a int, b int){
	numResult := a+b
	result <- numResult
}

// one goroutine 
func main() {

	ch := make(chan int)

	go processNum(ch) // one goroutine communicate to another goroutine 

	ch <- 10

	go sum(ch, 2,4) // call the another goroutine 

	res := <- ch
	fmt.Println(res) //received the data form another goroutine 
	time.Sleep(time.Second *2)

	// go func(){
	// 	ch <- 100
	// }()

	// value := <-ch

	// fmt.Print(value)
}