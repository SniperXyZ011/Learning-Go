package main

import "fmt"

//abhi tak ki khani, is humne ek payment stuct banya hai and uska ek method makePayment banay hai jo razorpay ka gateway use kr rha hai then finally razorpay ka hi pay method se payment execute kr rha hai, abhi tak koi problem nh, na hi abhi tak humne interface dekh hai, ab hume bola gya ki razorpay ki jgah humne stripe as a payment gateway use krna hai, so ab humne phle to stipe ke stucts and method define krne pdegai, not just that hume payment ke method ko bhi chnage krna pdega taki vo strip use kr ske, as you can see
// type payment struct{}

// func (p payment) makePayment(amount float32) {
// 	// razorpaygateway := razorpay{}
// 	stripepaygateway := strip{}// yha ab stripe krna pda
// 	// razorpaygateway.pay(amount) 
// 	stripepaygateway.pay(amount)//so ab hume stipe ka gateway bhi use krna pda
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	fmt.Println("Making payment using razorpay, total amount is", amount)
// }

// //so humne stripe ke stuct and methods define kr diye
// type strip struct{}
// func (s strip) pay(amount float32){
// 	fmt.Println("Making payment using stripe, total amount is", amount)
// }

//ye sab kaam krne mai, open close principal violate hua hai, which states ki classes should be open to extension but should be close to changes.

// so hume isko iss terha se approch krna pdega ki hume function mai chnages na krna pde,


//yha par humne dono payment gateway ek saath define kr diya and now hum jitna chaahe iska extensions kr skte hai, without chnaging a lot of code
// type payment struct{
// 	Gateway stripe
// }

// func (p payment) makePayment(amount float32){
// 	p.Gateway.pay(amount)
// }

// type razorpay struct{}
// func (r razorpay) pay(amount float32){
// 	fmt.Println("Making payment using razorpay", amount)
// }

// type stripe struct{}
// func (s stripe) pay(amount float32){
// 	fmt.Println("Making payment using stripe", amount)
// }

// now innn sab problem ko solve krne ke liye hamare pas aata hai interface ka concept, joki ek type of contract hota hai(ye typescript wale se toda diff hai) interface mai atleast ek method likhna mandatory hai, by convention interfaces ke names ke end mai "er" lag jata ha for ex: payement ka paymenter etc

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	Gateway paymenter //mast ab yha directly koi bhi gateway use krlo, bas we have to make sure ki usme koi pay method exist krt ho😉
}

func (p payment) makePayment(amount float32){
	p.Gateway.pay(amount)
}

type razorpay struct{}
func (r razorpay) pay(amount float32){
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}
func (s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe", amount)
}

//ab maan le ki tuje ni future paypal ka gateway add krna hai, then it's super easy
type paypal struct{}
func (p paypal) pay(amount float32){
	fmt.Println("Making payment using paypal", amount)
}


//ek key intresting point ki go mai interface ke andar jo hum method signature daalte hai and usko use krnte time agr hum koi struct use krte hai to hume explicitly mention nh krna hota (kucch lang mai "implement" keyword use hota hai), go ka compiler directly check kr leta hai ki jis struct ke liye ye interface use ho rha hai, usme koi function ka signature interface ke methods ke signature se match kr rha hai ya nh(super cool 🙃) 


//ab yha par hum apna koi fakepayment gateway create kr sakte hai, with a pay method taki hum isko test kr ske


func main() {
	totalAmt := 499.99
	// stripeGateway := stripe{}
	//ab agr mai yha par similarly razorpay ka method bhi likh du, and payment mai as a gateway pass krdu to ye kaam nh krgea since maine payment wale struct mai gate ko stripe define kr rkha hai therefore, muje payment struct ko bhi chnage krna pdega, mtlb this method is also prone to the chnages, ek problem ye hai ki testing mai problem ayegi since gateway hardcoded ho chuka hai as stripe
	paypalGateway := paypal{}
	newPayment := payment{
		Gateway: paypalGateway,
	}
	newPayment.makePayment(float32(totalAmt))
}