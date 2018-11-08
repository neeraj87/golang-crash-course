package main

import "fmt"

func main() {
	//ranges are used to loop through arrays, maps and slices

	ids := []string{"1", "2", "3"}

	for id := range ids {
		fmt.Printf("id is %s\n", id)
	}

	//looping through map
	//define an empty map
	emails := make(map[string]string) //make(map[key]value)

	//assign key value
	emails["Neeraj"] = "neeraj@email.com"
	emails["John"] = "john@email.com"
	emails["Deer"] = "deer@email.com"

	//Loop over map using keys
	//https://play.golang.org/p/JGZ7mN0-U-
	for k := range emails {
		fmt.Printf("key - %s and value - %s\n", k, emails[k])
	}

}
