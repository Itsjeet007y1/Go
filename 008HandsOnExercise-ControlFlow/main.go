package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(256)
	fmt.Printf("The value of x is %v\n", x)

	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x >= 256 {
		fmt.Println(("between 100 to 200"))
	} else {
		fmt.Println("More than 200")
	}

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}
