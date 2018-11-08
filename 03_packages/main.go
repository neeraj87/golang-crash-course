package main

import (
	"fmt"
	"math"

	"github.com/neeraj87/go_crash_course/03_packages/myownpackage"
)

func main() {
	fmt.Println("Demoing different packages")
	fmt.Println(math.Floor(2.8))
	fmt.Println(myownpackage.SayHello("neeraj"))
}
