package main

import (
	"fmt"
	"strings"
)

// also generics type any like this printSlice[T any], there any means any type of (int, string, bool, float)
// comparable is used to avoid hard coded define type like int | String | boole | float32 | float64 etc...

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}

}

// func printSlice[T int | string | bool | float32](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}

// }

type stack[T any] struct{
	element []T
}
func main() {
	nums := []int{1, 2, 3, 4}
	names := []string{"TP", "Tej", "Tarun"}

	printSlice(nums)
	printSlice(names)

	myStack := stack[string]{
		element: []string{"tej", "tp"},
	}
	fmt.Println(myStack)
}