package main

import (
	"fmt"
	// "maps"
)

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

	// now here we will learn how to check if a key - val pair exist in a map
	
	new_map := map[string]string{"name" : "Shyam", "student" : "True"}

	//now we want to check if name exists in this map
	v, ok := new_map["you"] // this syntx will return => value, <a boolean value>, isko hum usually (_, ok) likhte hai in go, _ jab hume value use nh krna, baki if you want to use the value as well then koi dusra naam likhan pdgea jise yha par v likha hai  
	if ok {
		fmt.Println("We found the value :", v)
	}else{
		fmt.Println("We didn't find the value")
	}

	// fmt.Println(maps.Equal(m1, m2)) to check if two maps are equal or not, just like slices mai tha

}