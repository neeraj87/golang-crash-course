package main

import "fmt"

var city = "Pune"

func main() {

	//creating variables

	//1. using var keyword (can be created outside the function)
	var name = "Neeraj"
	var age = 31 //int
	var isCool = true
	var distance = 5.6 //this is float

	fmt.Println(name, age, city, isCool, distance)

	//2. using shorthand method (we can only use this method inside a function)
	lastname := "Jadhav"
	fmt.Println(lastname)

}
