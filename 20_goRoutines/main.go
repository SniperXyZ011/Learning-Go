package main

import (
	"fmt"
	"time"
)

func task(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10;i++ {   
		// task(i)  //ye hai aam zindigi using a single core thread
		//now using go rounting to make it non blocking concurrent, mentos zindigi
		go task(i)
	}

	//yha par hume main func to toda rok kar rkhna pad rha hai so that ye func pura exit na ho jaygea
	time.Sleep(time.Second*2) //boom 
}