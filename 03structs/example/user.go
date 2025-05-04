package main  

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func isMinor(u User) bool{
	if (u.Age > 21){
		return true 
	}else{
		return false 
	}
}

func changeAge(u *User, age int){
	u.Age = age 
}

//reciever / method
//allows for user.setName()
func (user *User) setFirstName(name string){
	user.FirstName = name
}