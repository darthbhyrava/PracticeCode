package main

import ("fmt"
	"os"
	"log"
	"io/ioutil"	
	"strconv"
)

func main() {

	var v_int int = 23
	var v_float float64 = 23.4
	var v_string1 string = "100"
	var v_string2 string = "111.3"

	//Casting, converting from one type to another.
	fmt.Println(int(v_float))
	fmt.Println(float64(v_int))
	newint, _ := strconv.ParseInt(v_string1, 0, 64)
	newfloat, _ := strconv.ParseFloat(v_string2, 64)
	fmt.Println(newint)
	fmt.Println(newfloat)

	//Creating a file, with error handling. log.Fatal() prints error and exits.
	file, err := os.Create("samp.txt")
	if err!=nil{
		log.Fatal(err)
	}

	//Writing Data onto a String
	file.WriteString("This is some random text")
	file.Close()

	//Reading Data from a File 
	stream,err:=ioutil.ReadFile("samp.txt")
	if err!=nil{
		log.Fatal(err)
	}
	readString:=string(stream)
	fmt.Println(readString)


}