package main 

import "fmt"

func main() {

	//Arrays Declaration I
	var fnum[3] float64
	fnum[0]=12
	fnum[1]=1.23
	fnum[2]=8.27
	fmt.Println(fnum[2])
	//Arrays Declaration II
	fnum2 := [3]float64 {1,2,3}
	//Loops in Arrays
	for i, value := range fnum2 {
		fmt.Println(value,i)
	}
	for _, value := range fnum2 {
		fmt.Println(value)
	}

	//Slices
	numSlice := []int {5,6,7,8,9,10}
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[1]=",numSlice2[1])
	//Empty Slices: __slicename__ := make([]type, x|assign default 0 to x values initially, y|max size of slice is y)
	numSlice3 := make([]int, 6, 10)
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)
	//Appending elements to slices: __slicename1__ = append(__slicename2__, elements to be appended separated by commas)
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3)


	//Maps: _mapname_ := make(map[_keydatatype_]  _valuedatatype_)
	varMap := make(map[string] int)
	varMap["ageJohn"] = 12
	varMap["ageRam"] = 42
	fmt.Println(varMap,len(varMap))
	delete(varMap, "ageJohn")
	fmt.Println(varMap,len(varMap))

	





}