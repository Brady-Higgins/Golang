package main 

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main(){
	accounts := make(map[string]int)
	accounts["account1"] = 1000
	accounts["account2"] = 2000

	//deletes key & value
	delete(accounts,"account1")

	userDB := make(map[int]*User)
	userOne := User{FirstName: "andrew", LastName: "Schittz", Age: 21}
	userDB[0] = &userOne

}