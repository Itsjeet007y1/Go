package main

import "fmt"

func main() {
	x := [...]string{"Ram", "Shyam", "Mohan"} // array

	xs := []string{"Ram", "Shyam", "Mohan"} //slice
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(xs)

	// for range for slice
	for i, e := range xs {
		fmt.Printf("index: %v - value: %v\n", i, e)
	}

	// blank identifier to not use a returned value
	for _, e := range xs {
		fmt.Printf("value: %v\n", e)
	}

	xs = append(xs, "Raja", "Jitendra")

	fmt.Println(len(xs))
	fmt.Println(len(xs) > 0)
}
