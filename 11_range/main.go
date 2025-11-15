package main

import (
	"fmt"
)

//rage is used for the iteration over a data structue
func main(){
	//let's create a slice 
	my_slice := []int{1, 2, 3}
	//now let's iterate over this slice using for loop

	// for i := 0; i < len(my_slice); i++{  //ye hai normal zindigi
	// 	fmt.Println(my_slice[i])
	// }

	//now let's see the mentos zindigi using range keyword, range will return two things, idx, val
	//let's see in action

	for i, val := range my_slice{ //mast hai na, also if you don't need to use the idx just name it _, it will not cause any issue
		fmt.Println(i, val)
	}


	//abhi khtm nh hua, ab map's ko iterate krte hai
	my_map := map[string]string {"name" : "shyam", "city" : "delhi"}
	for key, val := range my_map{  //same concept use ho rha hai, baki agr convention follow krna ho to, (k,v), use kreo
		fmt.Println(key, "-> ", val)
	}

	//now now this isn't the end, we can travarse on string as well
	my_string := "Shyam Kumar"
	for i, c := range my_string{
		fmt.Println(i, c) // okay kucch craxy ho gya, isne char ki unicode(code point rune){ye kucch to hai jo aage pdgegai} value print kr di
		//ahmm also ye i idx nh hai, in technical terms isse starting byte of rune khte, aage pta chlega

		//now ab dekhte hai value kise print kre
		fmt.Println(string(c)) //now it works
	}	

	//tada😁 end of range
}