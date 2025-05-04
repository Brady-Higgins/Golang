package main

import "fmt"

func main() {
	avg := calculateAvg(1, 2, 3, 6, 8, 5)
	fmt.Println(avg)

	//runs when compiler gets to it 
	func(){
		fmt.Println("hey")
	}()

	//makes func callable within func 
	add := func(n1,n2 int) int {
		return n1 + n2
	}
	fmt.Println(add(1,2))

	//variables are now local
	value := 10
	addOne := func() int {
		value += 1
		return value 
	}
	fmt.Println(addOne())    // 11
	fmt.Println(addOne())    // 12
}

func calculateAvg(nums ...int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) /  float64(len(nums))
}