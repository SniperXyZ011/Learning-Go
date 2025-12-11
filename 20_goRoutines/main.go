package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(i int, w* sync.WaitGroup) {
	defer w.Done()  //yha par defer keyword ki used to tell ki ye kaam bilkul end mai krna hai
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10;i++ {   
		wg.Add(1)
		// task(i)  //ye hai aam zindigi using a single core thread
		//now using go rounting to make it non blocking concurrent, mentos zindigi
		go task(i, &wg)
	}

	wg.Wait()
	//yha par hume main func to toda rok kar rkhna pad rha hai so that ye func pura exit na ho jaygea
	// time.Sleep(time.Second*2) //boom 
	//we don't use sleep wala method in real use kuki hume nh pta kitne der rokna hai, therefore comes the concept of waitgroup, joki sync package se aata hai, isme 3 methods yaad rkhne hai, Add(), Done(), Wait()

}