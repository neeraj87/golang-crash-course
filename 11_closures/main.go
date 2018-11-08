package main

import "fmt"

//adder is a function which returns a function (this inner function takes an int and returns an int)
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//closures are anonymous functions, which we can define inline
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

}
