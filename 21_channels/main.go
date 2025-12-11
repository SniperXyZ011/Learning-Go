package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

//what is a channel? so basically ek channel ka use hota hai to pass data bw different go routines
//here we received data in thorugh channel
// func processNum(nums chan int){
// 	for num := range nums {
// 		fmt.Println(num)
// 		time.Sleep(time.Second)
// 	}
// }

func sum(result chan int, a, b int){
	calc := a + b
	result <- calc
}

func task(done chan bool){
	fmt.Println("Processing..")
	defer func ()  {
		done <- true //yha par data send kr diya on channel so ab ye block ho jaygea jab tak main mai isko recover nh krte	
	}()
}


func emailSender(emailChan <- chan string, done chan <- bool){ // to make it more type safe we can make channels as only send and only recive channels for send only example : (done chan <- bool) and for recive only : emailChan <- chan string
	defer func ()  {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("Email is begin sent to ", email)
		time.Sleep(time.Second)
	}
}

func main(){
	//channels hums make ka use krke bana skte hai, using chan keyword
	//see in action
	// messageChannel := make(chan string)
	// messageChannel <- "Shyam" //yha arrow indicate kar rha hai ki ye string iss channel mai daala jaa rha hai
	// msg := <- messageChannel   //note channels are blocking in nature(mtlb ki ye execution ko block krke rkhega jabtak dusre end se uss channel ka data use na ho)

	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100) //sending data
	// }

	//ab hum data receive kregai
	// result := make(chan int)
	// go sum(result, 4, 5)

	// res := <- result
	// fmt.Println(res) //yha par humne dusre routine ko kaam diya and uska result waps mangya to main routine

	//ab jo same kaam humne wait group ka use krke kiya tha that was ki main ko khtm hone se rokna, vo same kaam hume isse bhi kar skte hai
	//let's see that in action
	// done := make(chan bool) //ek channel bna liya (hum channel ke blocking nature ka use krne wale hai)
	// go task(done)
	// <- done //yha par humne vo data recover kr liya, now iske baad hi main func end hoga (as we can see ki humne uss data ka kucch use nh kiya)(agr iss line ko cmnt kar doge ko you'll se ki task wala routine end hone se phle hi main end ho gya)


	//abhi tak hum unbuffered channel use kr rhe hai, jo har ek i/o ke liye blocking hote hai, unlike buffered channel jo jabtk channel ki capacity bhr nh jati tab tak non blocking hota hai
	//let's see in action
	emainChan := make(chan string, 100)//buffered channel with size 100/
	done := make(chan bool)

	// emainChan <- "Shyam@gamil.com"
	// emainChan <- "Kumar@gmail.com"
    //originally in case of blocking channel ye code isske neeche kabhi paunch hi nh paata but since hamara channel buffered hai iss barr koi dikkat nh hogi
	// fmt.Println(<- emainChan)
	// fmt.Println(<- emainChan)

	go emailSender(emainChan, done)
	for i := 1; i <= 10; i++ {
		emainChan <- fmt.Sprintf("sam%d@gamil.com", i)
	}

	fmt.Println("Done sending")
	close(emainChan) //buffered channel mai channel close krna is imp wrna crash ho jayega
	<- done

}