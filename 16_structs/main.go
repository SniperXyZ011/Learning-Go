package main

import (
	"fmt"
	"time"
)

//go mai we don't have classes so we use structs

//chlo ek ecommerce mai order ke liye struct bante hai
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nano second tk ka precision hai iska
}

//ab yha ek aisa func create kro jo status ko change kre, stuct ke hisab se
func (o *order) changeStatus(status string){ //ye iska syntax thoda azeb hai but practice me aa jaygea, o likha caue vo struct ka first word hai(conventions baki teri marzi), also hume isko as a reference pass kiya hai func ke parameter mai kuki hume actual value ko change krna hai uski copy ko nh, also ab ques aane chiye ki fir o.status = status krte time bhi to *o.status = status krna chiye tha so yes aisa hona chiye acc to rule but struct hamari madat kr rha hai isme and isley hume * dobar nh likhna pdta so fundamentals abhi bhi vhi hai
	o.status = status
}

//now now ab zara hum thodi constructor ki baat kre, so go mai as such direct constructor nh hai, hume khud jugad lagna pdega constructor ke liye
//let's see this in action
//ab ek new product ka struct bante hai
type product struct{
	id string
	name string
	price float32
	quantity int
}

//ab stuct ke just baad hum ek func likhte hai isko initilize krke ke liye just like a constructor
func newProduct(id string, name string, price float32, quantity int) *product{ //as we can see humne values lekr newProduct bna diye and finally uska pointer return kr diya

	//yha par ab jo kucch bhi extra kaam krna hai before object creation vo kar skte ho
	product1 := product{
		id: id,
		name: name,
		price: price,
		quantity: quantity,
	}

	return &product1  //finally hume new product kr pointer return kregai
} 

//ek aur topic dekhte hai jisko khte hai struct embeddings jisme hum ek struct mai kisi dusre struct ko as a field rkhte hai
//jise maan le orders mai product as a field 
//let's take a new example
type student struct{
	rollNo int
	name string
	gender string
}

//ab hum bante hai a new struct called class jisme students honge
type class struct{
	classNo int
	section string
	student  //directly student wala struct as a field set kr diya
}

func main() {
	//normal syntax 
	// var newOrder = order {
	// 	id : "1",
	// 	amount : 30.44,
	// 	status : "pending",
	// 	createdAt: time.Now(),
	// }

	//shorthand 
	newOrder := order {
		id: "1",
		amount: 69.69,
		status: "received",
		createdAt: time.Now(),
	}

	fmt.Println(newOrder)
	fmt.Println(newOrder.status)

	newOrder.changeStatus("pending")
	fmt.Println(newOrder.status) //yha pending milega

	//let's use the product struct
	myProduct := newProduct("1", "guitar", 499.99, 56)
	fmt.Println(*myProduct) //(direct bhi print kr skte ho struct derefece ka kaam kr dega)



	// now now this isn't the end, agr thume bs ek baar liye koi struct bnana hai and use krna hai to uska bhi intezam hai
	newUser := struct { //mast syntax hai na abhi bano and turant use krlo, formally called inline struct
		name string
		age int
		gender string
	}{
		name: "shyam",
		age: 20,
		gender: "male",
	}

	fmt.Println(newUser)

	//chlo ab jara struct embedding ka example dekhte hai
	newStudent := student{
		name: "shyam",
		rollNo: 14,
		gender: "male",
	}

	newClass := class{
		classNo: 11,
		section: "A",
		student: newStudent,
	}

	fmt.Println(newClass)
}


//note: agr tum koi field ki value nh doge to usme uss type ki zerovalue aa jayegi