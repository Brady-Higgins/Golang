package main

import (
	"fmt"
	"sort"
)

func main() {
	// Array
	//0 based indexing
	var numsArray = []int{1, 2, 3, 4}
	fmt.Println(numsArray[2])

	// Slice
	var numSlice []int 
	numSlice = append(numSlice,1)
	numSlice = append(numSlice,2)
	numSlice = append(numSlice,3)

	sort.Ints(numsArray)

	for i:=0; i <len(numsArray); i++{
		fmt.Println(numsArray[i])
	}
	for index, val := range numsArray{
		fmt.Println("index: ", index, " value: ", val)
	}
	for _, val := range numsArray{
		fmt.Println("value: ", val)
	}
}