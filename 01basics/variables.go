package main

import "fmt"

// Globalvar := "not allowed"     
var globalPrivateVar = "global private"
var GlobalPublicVar = "global public"

func Vars(){
	//public
	Var1 := "short dec: in functions only, use 90% of the time"
	var Var2 = "full dec w/ inferred: outside functions mainly"
	var Var3 string = "full dec w/ explicit: when strict type should be enforced, lets compiler warn you"
	fmt.Println(Var1)
	fmt.Println(Var2)
	fmt.Println(Var3)
	//private
	var1 := "private short declaration"
	fmt.Println(var1)
	
	var tr bool = true
	fmt.Println(tr)
}
