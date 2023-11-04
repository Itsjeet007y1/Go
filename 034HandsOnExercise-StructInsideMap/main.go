package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "Jitendra",
		last:  "Kumar",
		favIC: []string{"Butter", "Choco", "Vanila"},
	}
	p2 := person{
		first: "Dharam",
		last:  "chacha",
		favIC: []string{"Butter", "Vanila"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, "favourite icecreams are: ")
	for _, e := range p1.favIC {
		fmt.Print(e, " ")
	}

	fmt.Println("-------------------------------------")

	// Struct inside a map
	m := map[string]person{
		p1.last: p1,
		p1.last: p1,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIC {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
