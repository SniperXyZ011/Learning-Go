package main


//for looping in go ther is no "while" or "do while" only "for loop" is availbe
func main(){
	// i := 1
	//implementing while loop using for keyword
	// for(i <= 10){
	// 	println(i)
	// 	i++
	// }

	//infinite loop
	// for{
	// 	println("infinte")
	// }

	//classic for loop
	for i := 1; i <= 10; i++ {
		println(i)
		if i == 7{
			i++
			continue
		}
		if i == 7{
			break
		}
	}

}