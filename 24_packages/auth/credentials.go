package auth

import "fmt"

// func loginWithCredentials(username string, password string) {  //ye func abhi locally scoped hai therefore iska use bas abhi issi folder mai ho skta hai, as we know ki js wagera mai hum export import jise concpet ka use krte the but in go we don't have to do all those complicated things it's much easier, hume bs function ko captial letter se start krna hota hai that's it
// 	fmt.Println("Logging user using username and password") 
// }

func LoginWithCredentials(username string, password string){
	//logic goes here 
	fmt.Println("Logging user via username and passoword ", username, password)
}