package main

import "fmt"

//this is passing argument by value, so iska actual main func wali value par asar nh hoga
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("Value in changeNum :", num)
// }

//now same func with passing by address
func changeByAddress(num* int){
	*num = 5
	fmt.Println("in func : ", *num)
}
func main() {
	num := 1
	// changeNum(num)
	changeByAddress(&num)
	fmt.Println("This is in main func :", num)
}


//crux of this part : agr varible as a refernce pass krna hai func mai then use &, and for the function use * to get the referece of that variable also use * again to derefrence