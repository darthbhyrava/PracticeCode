//My first Go program.

package main

//Equivalent of C's import<stdio.h>
import "fmt"

//Where magic happens. :P
func main() {
	
	//Customary Hello World Message:
	fmt.Println("Hello World")

	//Declaring some variables:
	var age int = 40
	var frac float64 = 3.14
	var myName = "Harsh" 
		
	//assignment requires the := operator
	randNum := len(myName)
	
	//Println() prints a new line by default, Printf() does not, unless \n is used. 
	fmt.Println(age,frac,randNum)
	fmt.Printf("%f %d \n", frac, age)

	//Some Printf() format specifiers
	fmt.Printf("%b %c %x %e %T \n", 100, 44, 17, frac, frac)

	//Logical Operators
	fmt.Println(" true && false = ", true && false)
	fmt.Println(" true || false = ", true || false)
	fmt.Println(" !true = ", !true)

	//Loops
	i := 1
	for i<=10 {
		fmt.Println(i)
		i++
	}
	for j:=1; j<5; j++ {
		fmt.Println(j)
	}

	//Branching
		//if
	if age<20 {
		fmt.Println("Teenager")
	} else if age <40 {
		fmt.Println("MidLife-Crisis. :P")
	} else {
		fmt.Println("Gandalf!")
	}
	//Switch
	heyJude :=2

	switch heyJude {
		case 1: fmt.Println("Hey, Jude")
		case 2: fmt.Println("Don't be so sad.")
		default: fmt.Println("Beatles!")
	}
}