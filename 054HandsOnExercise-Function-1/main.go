package main

import "fmt"

func main() {
	x := sum([]int{1, 2, 3, 45, 5, 44, 3})
	fmt.Println(x)
}

func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}
