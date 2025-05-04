package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	personOne := Person{
		FirstName: "Sherlock",
		LastName:  "Holmes",
		Age:       40,
	}
	fmt.Println(personOne.Age)
}