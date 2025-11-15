package main

import "fmt"

//ab finally hum seekhegai, functions ke baare mai

func my_first_function(){
	fmt.Println("Yay, my function called")
}

//now let's make a add function
func add(a int, b int) int { //hume  parameter ka type likha as well as function ky return kr rha hai vo bhi likha just like typescript
	return a + b  // ek aur mazadar baat, agr parameter mai ek hi type ke variables ho to baar bar type likhne ji zarurat nh hai, ek baar hi likh skte hai jise iss case mai (a, b int)
}

//go mai a function can return more than one values
func more_func()(string,string, string){ //yha par () dena imp hai
	return "one", "two", "three"
}


//go mai function 1st class citizen hote hai, kucch to hota hai iska matlab, like func can be assigned to a var, another func etc
//let's see this in action

func add_new_fun(fn func(a, b int) int, fn2 func(a, b int) int, x int, y int, z int, w int) int{
	//okay so isme hum apne add function ko daalegai and final result nikalegai
	firstVal := fn(x, y)
	secondVal := fn2(w, z)
	return firstVal + secondVal
}
func main(){
	my_first_function()

	fmt.Println(add(1, 2))

	// fmt.Println(more_func())
	num1, num2, num3 := more_func()
	fmt.Println(num1, num2, num3)

	//now let's see func in func 
	fmt.Println(add_new_fun(add, add, 1, 2, 3, 4))
}