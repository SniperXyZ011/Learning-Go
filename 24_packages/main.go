package main

import (
	"fmt"

	"github.com/SniperXyZ011/Learning-Go/auth"
	"github.com/SniperXyZ011/Learning-Go/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("Shyam", "1233")
	fmt.Println(auth.GetSession())

	newUser := user.User{
		Email: "hell@gmail.com",
		Name: "Shyam",
	}

	fmt.Println(newUser)
	//ab zara color mai print krte hai
	color.Cyan(newUser.Email)
}

//final words : export ya expose krne ke liye hume Capital letter se shuru kro, for privite ke liye use small letters

//now ab i am trying to use external pakages halaki go mai almost saari cheezo ke liye packges inbuilt hai magar tab bhi agr use kna hai than go on : pkg.go.dev wha par koi bhi pakage ko search kro and then copy kro, it would look like this : 	github.com/fatih/color
// ab terminal par jakr run this : go get github.com/fatih/color 
//that's it and you can then use it

// note : go mod tidy -> ye cmd cheeze fix krne + linting issue ko fix krta hai(very powerful cmd)