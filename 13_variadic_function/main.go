package main

import "fmt"

//let's create our own variadic function
func sum(num ...int)int{ // yha par parameters mai we only want int, but agr any kar diya then it will be same as println
	//num ek slice ban jaegi
	res := 0
	for _, v := range num{
		res += v
	}

	return res
}


func main() {
	//let's see what is this variadic function,
	fmt.Println(1, 2, 3, "hello") //look at this println func this is taking n number of arguments right vo bhi kisi bhi type ki values, this is variadic function

	//calling my own varidic function
	fmt.Println(sum(1, 2, 3, 4, 5))

	//ek aur crazy cheez, maan le tuje kud ye 1,2,3,4 nh likhna very boring then you can directly pass a slice
	//let's see how
	my_par := []int {1, 2, 3, 4, 5}
	fmt.Println(sum(my_par...)) //dekha spread operator jise hi hai 

}