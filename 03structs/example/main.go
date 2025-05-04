package main

import "fmt"

func main() {
	u := User{
		FirstName: "holly",
		LastName:  "Pot",
		Age:       21,
	}
	if isMinor(u) {
		fmt.Println("Yup")
	}else{
		changeAge(&u,12)
	}
}