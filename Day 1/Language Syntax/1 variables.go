package main

import (
	"fmt"
	"math"
)

func main() {
	//exercise 1
	var a int
	var b string
	var c bool
	
	fmt.Println(a,b,c)

	// exercise 2
	a=1
	b="strung"
	c=true

	fmt.Println(a,b,c)

	//exercise 3
	var pi float32
	pi = math.Pi 

	fmt.Printf("pi(%T) = %v",pi,pi)
}
