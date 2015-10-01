package main 
import "fmt"

//func funcname(parameters) returntype {funcdesc}
func add(list []int) int{
	sum:=0
	for _,val:=range list {
		sum+=val
	}
	return sum
}

//functions in Go can return multiple values
func nextnums(num int) (int, int) {
	return num+1, num+2
}

//for functions where you don't know number of parameters
func addm(vals ...int) int {
	sum:=0
	for _, value:=range vals{
		sum+=value
	}
	return sum
}

//Recursion
func fact(n int) int{
	if n==0 {
		return 1
	} else {
		return n*fact(n-1)
	}
}

//Recover Function
func safeDiv(i1, i2 int) int{
	defer func(){
		fmt.Println(recover())
		//This will thro up error but still continue program. To continue without printing error, just use recover().
	}()
	sol:=i1/i2
	return sol
}

func main() {

	numlist:=[]int {1,2,3,4,5}
	fmt.Println("Sum =", add(numlist))	
	n1,n2:=nextnums(3)
	fmt.Println(n1,n2)
	fmt.Println("Sum =", addm(1,2,3,4,5))

	//defer executes adjoining function after the parent function has ended
	defer fmt.Println(add(numlist))

	//Demonstration of the recover function
	fmt.Println(safeDiv(2,0))
	fmt.Println(safeDiv(2,1))

	//Defining functions inside other functions.
	n3:=3
	doub:= func() int {
		n3*=2
		return n3
	}
	fmt.Println(doub())
	fmt.Println(fact(5))
}

