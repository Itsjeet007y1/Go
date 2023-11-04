package main

import "fmt"

func main() {

	jk := []string{"Jitendra", "Kumar", "Choco", "Vanilla"}
	sk := []string{"Santosh", "Kumar", "Harry", "Potter"}

	fmt.Println(jk)
	fmt.Println(sk)

	members := [][]string{jk, sk}

	fmt.Println(members)

	fmt.Println("--------------------------------")

	for i, v := range members {
		fmt.Println(v, i)
		for a, b := range v {
			fmt.Println(b, a)
		}
	}
}
