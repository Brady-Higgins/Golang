package main

import "fmt"

func main() {
	var n int = 2
	fmt.Println("int is: ", n)
	changeInt(&n,4)
	fmt.Println("int is: ", n)
}

func changeInt(old_n *int, new_n int) {
	fmt.Println("pointer address in hex",old_n)
	*old_n = new_n
}