package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46}
	for i, v := range xi {
		fmt.Printf("index %v \t value %v\n", i, v)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
		fmt.Println()
	}

	m := map[string]int{
		"James":      42,
		"MoneyPenny": 32,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
