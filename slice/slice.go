package main

import (
	"fmt"
	"slices"
)

func main() {
	// var nums []int
	// fmt.Print(nums)

	// var nums = make([]int,0)
	// // fmt.Print(nums)

	// nums = append(nums, 1)
	// nums = append(nums, 3)
	// fmt.Print(nums," ")
	// fmt.Println(len(nums))
	// fmt.Print(cap(nums))

	// nums := []int{}
	// // fmt.Print(nums)
	// nums = append(nums, 1)
	// nums = append(nums, 3)
	// fmt.Println(nums)

	var nums = []int{1,2,3}
	var nums1 = []int{1,2,3}
	fmt.Println(slices.Equal(nums, nums1))

	//2D slice
	var nums2 = [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(nums2)
}