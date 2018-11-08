package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

//Value receiver method for a struct
//a value reciever method just takes values from a struct and just display
func (myperson Person) greeting() string {
	return "Hi, my name is " + myperson.firstname + " " + myperson.lastname + " and my age is " + strconv.Itoa(myperson.age)
}

//Pointer receiver method for a struct
//a pointer receiver method reads struct value and changes them
func (myperson *Person) hasBirthday() {
	myperson.age++
}

func (myperson *Person) changeLastName(newlastname string) {
	myperson.lastname = newlastname
}

//SINCE I HAVE BEEN USING myperson AS THE RECEIVER NAME FOR REPRESENTING PERSON IN ALL THE METHODS, I HAVE TO KEEP USING IT
//IF I USE ANY NEW RECEIVER NAME LIKE (person *Person) ITS GONNA THROW AN ERROR

func main() {
	//init person using struct
	person1 := Person{
		firstname: "John",
		lastname:  "Doe",
		city:      "Austin",
		gender:    "M",
		age:       33}

	fmt.Println(person1) //{John Doe Austin M 33}

	//or we can initilize a struct from variables
	fname := "Jane"
	lname := "Doe"
	city := "Austin"
	gender := "F"
	age := 44

	person2 := Person{fname, lname, city, gender, age}

	fmt.Println(person2) //{Jane Doe Austin F 44}

	//get single field
	fmt.Println(person1.firstname)
	fmt.Println(person2.age)

	//chaning a field ina struct
	person1.city = "Boston"
	fmt.Println(person1) //{John Doe Boston M 33}

	//calling pointer receiver method
	person1.hasBirthday() //increases the age by one

	//calling a struct method
	fmt.Println(person1.greeting())

	person2.changeLastName("Lynch")
	fmt.Println(person2)

}
