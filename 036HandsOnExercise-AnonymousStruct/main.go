package main

import "fmt"

func main() {
	p1 := struct {
		first   string
		friends map[string]int
		drints  []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny":  27,
			"Q":      87,
			"Ramesh": 44,
		},
		drints: []string{
			"Water",
			"Milk",
		},
	}

	for k, v1 := range p1.friends {
		fmt.Println(p1.first, "- friends", k, v1)
	}

	for _, v2 := range p1.drints {
		fmt.Println(p1.first, "- drinks", v2)
	}
}
