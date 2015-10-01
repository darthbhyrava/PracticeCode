package main 

import( "fmt"
	"time"
	"strconv"
)

var pizzanum=0
var pizzaname=""

//Channels - to pass data between go routines. 
//makeDough receives a channel as input, and creates a new "pizza"
func makeDough(stringChan chan string) {
	pizzanum++
	pizzaname="Pizza #"+strconv.Itoa(pizzanum)
	fmt.Println("Make Dough and send for Sauce")
	stringChan <- pizzaname
	time.Sleep(time.Millisecond*10)
}
func addSauce(stringChan chan string) {
	pizza:= <- stringChan
	fmt.Println("Add Sauce and send",pizza,"for toppings")
	stringChan <- pizzaname
	time.Sleep(time.Millisecond*10)
}
func addToppings(stringChan chan string) {
	pizza:= <- stringChan
	fmt.Println("Add Toppings to",pizza,"and ship.")
	time.Sleep(time.Millisecond*10)
}

func main() {
	//Create a channel which can hold a string; it will pass a string value between all of our go routines. 
	stringChan:=make(chan string)
	
	//We'll keep creating pizzas till 3 pizzas are made
	for i:=0;i<3;i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond*5000)
	}

}