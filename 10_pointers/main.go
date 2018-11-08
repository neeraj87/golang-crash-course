package main

import "fmt"

func main() {
	//EVERYTHING IN GOLANG IS PASS BY VALUE

	//pointer allows you to point location of a value in the memory
	//we can use pointers to pass a lot of data stored at address and we can pass address to increase efficiency
	a := 5
	b := &a //so b is pointer of a

	fmt.Println(a)
	fmt.Println(b) //gives memory address of where a is stored in the system

	//check type
	fmt.Printf("data type of a - %T \n", a)
	fmt.Printf("data type of b - %T \n", b)

	//reading value from address
	fmt.Println(*b)
	//OR
	fmt.Println(*&a) //because b is set to &a

	//changing the value with pointer
	*b = 10
	fmt.Println(a) //we get 10
}
