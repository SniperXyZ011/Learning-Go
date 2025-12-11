package main

import "fmt"

//generics were not present in go till version 18
func printIntSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string){
	for _, item := range items {
		fmt.Println(item)
	}
}

//yha par hume same kaam ke liye dobara faltu code likhna pd rha hai jaki bs datatype hi different hai, here comes the concept of generics

func printSlice [T any](items []T){ //isko agr scoped krna ho, i.e abhi ye any type ke liye kaam kr rha hai but if we want ki ye only kucch datatype ke liye hi kaam kre then we can define it like this 😎 : func printSlice [T int | string] (items []T)
	for _, item := range items {
		fmt.Println(item)
	}
}


//ab same cheez hum struct mai kise kr skte hai
//ek stack ka struct bnate hai
type stack[T any] struct {
	elements []T
}

func main() {
	// printIntSlice([]int {1, 2, 3, 4})
	// printStringSlice([]string {"Shaym", "Anubhav", "Aryan"})

	//works for same func
	printSlice([]int {1, 2, 3, 4})
	printSlice([]string {"Shaym", "Anubhav", "Aryan"})

	myStack := stack[int]{
		elements : []int {1, 3, 4},
	}

	fmt.Println(myStack)
}