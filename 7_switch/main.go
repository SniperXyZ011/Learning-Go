package main

import (
	"fmt"
	// "time"
)

func main() {
	//simple switch
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("i is one")
	// case 2:
	// 	fmt.Println("i is two")
	// case 3:
	// 	fmt.Println("i is thee")
	// default:
	// 	fmt.Println("Na munna ye na ho payega")
	// }

	//multiple condtion switch
	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday :
	// 	fmt.Println("Maze hi maze chutty hai")
	// default :
	// 	fmt.Println("Majduri krna hoga lala")
	// }

	//type switch -> pta ni kucch fnuction wagera par lagega
	whoAmI := func(i interface{}){
		switch t := i.(type) {
		case int :
			fmt.Println("Data type of is INT")
		case float32 :
			fmt.Println("Data type is float")
		case string :
			fmt.Println("Data is string")
		case bool :
			fmt.Println("Data is boolean")
		default :
			fmt.Println("Ni pta bro ky de raha hai tu", t)
		}
	}

	whoAmI(3432234.234)

}
