package main

import ( "fmt"
	"math"
	"strings"
	"sort"
	)

//Taking parameters as pointers.
func poin(x *int) {
	*x=2
}

//Defining two structs in golang, and ensuing functions for the struct.
type Rectangle struct {
	height float64
	width float64
}
type Circle struct {
	radius float64
}
func (rect Rectangle) area() float64{
	return rect.height*rect.width
}
func (c Circle) area() float64{
	return math.Pi*math.Pow(c.radius, 2)
}


//Defining an interface
type Shape interface {
	area() float64
}
func getArea(shape Shape) float64{
	return shape.area()
}

func main() {	

	//Strings
	str:="Hello World"
	fmt.Println(strings.Contains(str,"lo"))
	fmt.Println(strings.Index(str,"lo"))
	fmt.Println(strings.Count(str,"l"))
	fmt.Println(strings.Replace(str,"l","x", 2))
	//Splitting Strings
	csvStr:="De Gea, Darmian, Smalling, Blind, Shaw"
	fmt.Println(strings.Split(csvStr,","))
	//Joining Strings
	jstr:=strings.Join([]string {"a","c","b"}, ":")
	fmt.Println(jstr)
	//Sorting Strings
	randLet:=[]string {"c", "a", "b"}
	sort.Strings(randLet)
	fmt.Println("Letters=", randLet)


	//Pointers
	x:=0
	poin(&x)
	fmt.Println("Value of x: ", x)
	fmt.Println("Memory Address of x: ", &x)

	//Structs
	rect:=Rectangle{height:10, width:5} //Or just go Rectangle{10,5} if you know the order. 
	fmt.Println(rect.area())
	circ:=Circle{30}
	fmt.Println(circ.area())

	//Using the interface
	fmt.Println("Circle and Rectangle Area =", getArea(circ), getArea(rect))
}
