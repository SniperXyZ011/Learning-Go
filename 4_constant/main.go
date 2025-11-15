package main

import "fmt"

const age = 21

func main() {
	const name string = "Shyam" //can't change this
	//grouping can be done for consts
	const (
		port = 300
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(port, "->", host)

}