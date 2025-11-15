package main

import "fmt"
func main() {
	// age := 21
	// if age > 18 {
	// 	println("You are eligible for voting")
	// } else {
	// 	println("You are not eligible for voting")
	// }

	//cool syntax 😎
	if age := 21; age > 18 {
		fmt.Println("Yes you are eligible since your are is :", age)
	}

	//sad part go dosen't have ternery operator so you have to use if-else always😔
}
