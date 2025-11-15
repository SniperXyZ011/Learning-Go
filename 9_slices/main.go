//basically hume go mai vectors chiye jo apna size dynamically increase kr sake so we have slice in Go

package main

import (
	"fmt"
	"slices"
)

//slice dynamic
//most useful construct in Go
//many inbuilt methods
func main() {
	//uninitilzed slice is nil(go mai nil khte hai baki lang mai null)
	// var num []int
	// num = append(num, 1)
	// println(num[0])
	// println(len(num)) // returns 0

	//agr tuje chiye ki slice ka size nil na ho kucch min size ho then use make method
	// var arr = make([]int, 2, 5) // agr tuje capacity already fix krni ho then you can put a third argument which will make it's capacity
	// println(len(arr)) //returns 2

	//capacity is max element which can be insterted, we have cap() method to find that
	// println(cap(arr)) //5
	// println(cap(num)) // 0

	//shorthand 
	// arr2 := []int{}

	//ek experiment
	// var meraArray []int
	// var meraArray2 []int
	// meraArray2 = append(meraArray, 1)
	// println("Mera array1 : ", meraArray == nil)
	// println("Mera array2 : ", meraArray2[0])


	//copy method to copy a slice 
	//copy(kisme krna hai, kissse krna hai)
	// arr1 := make([]int, 0, 5)
	// arr1 = append(arr1, 1, 2)
	// arr2 := make([]int, len(arr1))
	// copy(arr2, arr1)

	// fmt.Println(arr1, arr2)


	//slice operator
	// arr := []int{1, 2, 3, 4}
	// fmt.Println(arr[0:2]) //last wala index excluded hota hai

	//slice comparision
	arr1 := []int {1, 2, 3}
	arr2 := []int {1, 2, 3}

	fmt.Println(slices.Equal(arr1, arr2))
}
