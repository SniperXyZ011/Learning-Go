package main

import (
	"fmt"
)

//numbered sequence of specific length
func main(){
	var arr[4] string
	// for i := range(4){
	// 	arr[i] = i
	// }
	
	// for i := range(4){
	// 	fmt.Println(arr[i])
	// }
	// fmt.Println(len(arr))

	//now when we initilize an arry of type int  -> all 0's, for bool -> all false, for string -> empty string
	fmt.Println(arr)

	//direct values 
	num :=[4]int {1, 2, 3, 4}
	fmt.Println(num)


	//2d array
	nums := [2][2] int{{1,2},{1,2}}
	fmt.Println(nums)
}