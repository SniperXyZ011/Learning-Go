package main

import "fmt"

func main() {
	//how to create a map

	// mp := make(map[string]string) // make(map<keyword>[<type of key>]<type of value>)
	var mp = make(map[string]string)

	//set an element {key : val}
	mp["Name"] = "Shyam Kumar"

	//print element
	fmt.Println(mp["Name"])

	//lenght of map
	fmt.Println(len(mp))

	//how to delete an element use delete() method
	delete(mp, "Name") // format dyan rkhna
	fmt.Println(mp)

	mp1 := make(map[int]int)
	mp1[1] = 1
	mp1[2] = 1
	mp1[3] = 1

	fmt.Println(mp1)
	//clear method
	clear(mp1)
	fmt.Println("Cleared map1", mp1)


	//let's try to access a element which dosen't exist
	fmt.Println("Here acessing value of Hello keyword :",mp["Hello"]) //return an zerovalue(i.e agr type of value int hai so 0, for bool -> fasle, for string empty string) value (key point doesn't give an error)

	//prtical 
	mp2 := make(map[int]int)
	fmt.Println(mp2[9]) //returns 0

	//shorthand for map
	my_new_map := map[int]int {1: 4, 2 : 3}

	fmt.Println(my_new_map[1])

}