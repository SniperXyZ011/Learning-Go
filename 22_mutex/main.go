package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex //so now ab hum mutex use kregai
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func ()  {
		wg.Done()
		p.mu.Unlock()	
	}()
	p.mu.Lock() // yha lock kr diya
	p.views += 1 //yha modify hua
	// p.mu.Unlock() // yha unlock ho gya // now it will never run into race condtion // usually isko defer func mai daalte hai so that ki agr during func exection koi problem ho to bhi ye unlock ho jayega
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}

	// myPost.inc()
	// myPost.inc()// abhi 2 bar add hua no problem
	//let say ki 100+ views aa gya
	// for i := 0; i < 100; i++ {
	// 	myPost.inc() //yha bhi koi problem nh hogi kuki ye code abhi synchronously chal rha hai, but real lyf mai aisa nh hota
	// }

	//yha par hum routins ka use krke concurrently simulate krne ki kosish kr rhe hai 
	for i := 0; i < 100; i++ {
		wg.Add(1)
		myPost.inc(&wg) //hum ye dekh skte hai ki har baar zaruri nh hai ki final value 100 aaye, so this problem will be solved using mutex
	}
	wg.Wait()
	fmt.Println(myPost.views)
}