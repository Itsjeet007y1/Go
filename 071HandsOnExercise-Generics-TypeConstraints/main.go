package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 4))
	fmt.Println(addF(1.4, 4.3))

	fmt.Println(addT(1, 4))
	fmt.Println(addT(1.3, 4.4))
	fmt.Println("Jitendra", "Kumar")
}
