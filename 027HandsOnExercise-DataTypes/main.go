package main

import "fmt"

func main() {

	// create a slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 52, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59}
	// x = append(x, y[0])
	x = append(x, y...)
	fmt.Println(x)

}
