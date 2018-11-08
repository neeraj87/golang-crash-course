package main

import "fmt"

func main() {

	//for loop - long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//for loop - short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d \n", i)
	}

	//FizzBuzz challenge
	//loop through 1 to 100
	//if the number is divisible by 3, print Fizz
	//if the number is divisible by 5, print Buzz
	//if the number is divisible by both 3 and 5, print FizzBuzz

	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%15 == 0 { //to find if a number is divisible by both 3 and 5, we see if its divisible by 15 (product of 3 and 5)
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
