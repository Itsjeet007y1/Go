package main

import (
	"fmt"
	"math/rand"
)

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

	fmt.Println("------------------------------------")
	m := map[string]int{
		"James":      42,
		"MoneyPenny": 32,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("------------------------------------")
	age1 := m["James"]
	fmt.Println("The age of Bond", age1)

	if v, ok := m["James"]; ok {
		fmt.Println("There is James and bond's age is", v)
	}

	age2 := m["Q"]
	fmt.Println("The age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q and there is the zero value of int", v)
	}

	fmt.Println("------------------------------------")
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v\t x is %v\n", i, x)
		}
	}
}
