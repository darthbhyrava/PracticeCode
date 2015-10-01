package main 

import( "fmt"
	"time"
)

//This fn. will pause some events using time.Sleep(), and then we will use go routines which will help us run multiple events in parallel 
func count(id int) {
	for i:=0;i<10;i++ {
		fmt.Println(id,":",i)
		time.Sleep(time.Millisecond*1000)
	}
}

func main() {
	for i:=0;i<10;i++ {
		//executes go routine
		go count(i)
	}	
	//Another Sleep, so that this pause is long enough for all go routines to execute
	time.Sleep(time.Millisecond*11000)
}