package main

import "fmt"

func modify(a [5]int)  {
	a[0] = 6
}

func modify1(a *[5]int)  {
	a[0] = 6
}

func TwoDArray()  {
	var matrix [2][3]int = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
for  i:= 0; i < len(matrix); i++ {
	for j := 0; j < len(matrix[i]); j++ {
		fmt.Print(matrix[i][j]," ")
	}
	println()
	}
}


func matrix1()  {
	var matrix [2][3]int = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
for _, row := range matrix{
	for _, v := range row {
		fmt.Print(v," ")
	}
	fmt.Println()
}
}

func compareValueOrArray()  {
	arr := [3]int{1,2,3}
	arr1 := [3]int{4,2,6}
	fmt.Println(arr == arr1)

	for i := 0; i < len(arr); i++ {
		if(arr[i] == arr1[i]){
			fmt.Println("equal value")
		}else{
			fmt.Println("not equal value")
		}
	}
}
func main()  {
	var array [4]int
	var array1 [3]int
	array1[0] = 10
	fmt.Println(array1[0])
	array1[0]= 20
	fmt.Println("update the value of index zero ",array1[0])
	array = [4]int{12,13,14,15}
	fmt.Println(array)

	arr := [5]int {1,2,3,4,5}
	fmt.Println(arr)

	array2:=[...]int{1,2,3,4,5,6,7,8}
	fmt.Println(array2)
	fmt.Println("length of array Two is: ",len(array2))

	fmt.Println("Looping throw the array and take a operation!")

	for i := 0; i < len(array2); i++ {
		fmt.Print(array2[i]*2," ")
	}

	fmt.Println("Using range!")
	for index, value := range arr{
		fmt.Println(index, value)
	}
	fmt.Println("travese the array using range method without index")
	for _, value := range arr{
		fmt.Print(value," ");
	}

	for i := range arr{
		fmt.Print(i)
	}
	modify((arr))
	fmt.Println("After the modify ",arr)

	modify1((&arr))
	fmt.Println("afte the modify uisng pointer ",arr)

	TwoDArray()
	matrix1()

	compareValueOrArray()
}