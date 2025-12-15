package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {

	//some basics 

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error hai sir", err)
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	fmt.Println("Error hai sir", err)
	// 	panic(err)
	// }

	// fmt.Println("FIle name :",fileInfo.Name())
	// fmt.Println("File or folder :", fileInfo.IsDir())
	// fmt.Println("FIle size :",fileInfo.Size())
	// fmt.Println("FIle permissions :",fileInfo.Mode())
	// fmt.Println("FIle name :",fileInfo.ModTime())


	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//how to read a file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error found ", err)
	// 	panic(err)
	// }

	// defer f.Close() //file close krna is imp

	// buff := make([]byte, 25)
	// _, e := f.Read(buff)

	// if e != nil {
	// 	panic(e)
	// }

	// println("Data :", string(buff))

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//more simpler way to read file
	// f, err := os.ReadFile("example.txt") //yha without file open kiye read kr liya, it's good to use it for small files but don't use it for smaller files, since ye puri file ek saath memory mai load krdeta hai, so agr badi file hui to apne system ke pakode lag skte hai

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//read folders
	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()
	// fileInfo, err := dir.ReadDir(-1)
	// for _, data := range fileInfo {
	// 	fmt.Println(data.Name(), data.IsDir())
	// }


	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//creating files
	// f, err := os.Create("exmaple2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi i am shyam kumar, nice to meet you")
	// fmt.Println(f)

	// //writing using a slice of byte 
	// bytes := []byte("Hello everyone")
	// f.Write(bytes)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//streaming of data example, read from one file and write on another

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	
	// defer sourceFile.Close()
	
	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, e := reader.ReadByte()
	// 	if e != nil {
	// 		if e.Error() != "EOF" {
	// 			panic(e)
	// 		}
	// 		break
	// 	}

	// 	err := writer.WriteByte(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// writer.Flush()
	// fmt.Println("Done streaming....")


	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//how to delete a file
	
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("removed example2.txt successfully")

	
}