package main

import "fmt"

func main() {
	//define an empty map
	emails := make(map[string]string) //make(map[key]value)

	//define and initialize a map
	m := map[string]string{"key1": "val1", "key2": "val2"}

	//assign key value
	emails["Neeraj"] = "neeraj@email.com"
	emails["John"] = "john@email.com"
	emails["Deer"] = "deer@email.com"

	fmt.Println(emails)
	fmt.Println(m)
	fmt.Println(emails["Deer"]) //get a value by its key

	//Loop over map using keys
	//https://play.golang.org/p/JGZ7mN0-U-
	for k := range emails {
		fmt.Printf("key - %s and value - %s\n", k, emails[k])
	}

	//deleting from map
	delete(emails, "Deer")
	fmt.Println(emails)
}
