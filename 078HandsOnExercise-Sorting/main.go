package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 4, 2, 6, 77, 33, 24, 1, 43}
	xs := []string{"Ram", "Shyam", "Rahul", "Jitendra", "Q", "M", "MJ", "Rocky"}

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

	// sorting in reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	sort.Sort(sort.Reverse(sort.StringSlice(xs)))

	fmt.Println("Reverse sorting---------")
	fmt.Println(xi)
	fmt.Println(xs)
}
