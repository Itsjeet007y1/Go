package main

import "fmt"

func intDelta(a *int) {
	*a = 43
}

func sliceData(x []int) {
	x[3] = 99
}

func mapDelta(m map[string]int, s string) {
	m[s] = 43
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	sliceData(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 42
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])
}
