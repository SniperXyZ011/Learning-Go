package main

import "fmt"

//enumerated types
//go mai enum directly available nh hai, hum const + type keyword ka use krke isko implement krte hai

//note type keyword ka use hum custom type bannae ke liye bhi kar skte hai
type OrderStatus int

const (
	Recived OrderStatus = iota //yha par hum iota issley likh rhe hai kuki ye ek terika ka auto incremneter ke terha kaam kar rha hai, that is Recived ko 0 assign hua hai and next ko 1 assign hoga and so on..
	Confirmed //1 
	Delayed //2
	Processing //3
	Dispatched //4
)

//int ki jagha hum koi dusra datatype bhi use kr skta hai but then iota use nh hoga kud hume uska value likhna pdgea
func changeStatus(status OrderStatus){
	fmt.Println("Chnaging order status to ->", status)
}

func main() {
	changeStatus(Dispatched)
}
