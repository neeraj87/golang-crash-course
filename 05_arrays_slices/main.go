package main

import "fmt"

func main() {
	//with Go the Arrays have to be of fixed length and of fixed type
	var fruitArray [2]string

	//assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Banana"

	//declare and assign at same time
	vegArray := [2]string{"Tomato", "Cabbage"}

	fmt.Println(fruitArray)
	fmt.Println(vegArray)

	//SLICES - they are arrays with undetermined length
	fruitSlice := []string{"Mango", "Pineapple", "Watermelon"}
	fmt.Println(fruitSlice)

	//counting length of array or slice
	fmt.Println(len(fruitSlice))
}
