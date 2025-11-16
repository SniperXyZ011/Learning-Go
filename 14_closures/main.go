package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	incrementer := counter()
	fmt.Println(incrementer())
	fmt.Println(incrementer())
	fmt.Println(incrementer())

	//so let's see ki actual mai ho ky rha hai and wtf is excatly closures, so jo counter func likha hai humne vo ek anomyous func return kr rha hai jo value increment kr rha hai, what we know is ki agr koi func call hota hai then after execution vo call stack ke hat jata hai, i.e usne jo bhi memory use ki hai, for variable and everything vo ktm ho jati hai, but iss case mai agr hum dekhe to hum anomyous func ko outer func ka variable de rhe hai i.e vo varible uss func ke scope ke bhar ka hai, issi wjh se jab hume incrementer func ko baar baar call kr rhe hai to counter ki value increment ho rhi hai na ki shuru se start ho rhi hai, and this is happening due to the closure in Go, i know thoda heavy ho gya but easy hai give it some time.
}