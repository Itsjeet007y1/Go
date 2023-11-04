package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x and y are %v and %v\n", x, y)

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	// if x < 4 && y < 4 {
	// 	fmt.Println("Both are less than 4.")
	// } else if x > 6 && y > 6 {
	// 	fmt.Println("Both are greater than 6.")
	// } else if x >= 4 && y <= 6 {
	// 	fmt.Println("X is greater or equal to 4 and y is less than or equal to 6.")
	// } else if y != 5 {
	// 	fmt.Println("X is not equal to 5")
	// } else {
	// 	fmt.Println("Some other values.")
	// }

	for i := 0; i < 100; i++ {
		// fmt.Println(i)
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x and y are %v and %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than 4.")
		case x > 6 && y > 6:
			fmt.Println("Both are greater than 6.")
		case x >= 4 && y <= 6:
			fmt.Println("X is greater or equal to 4 and y is less than or equal to 6.")
		case y != 5:
			fmt.Println("X is not equal to 5")
		default:
			fmt.Println("Some other values.")
		}
	}
}
