package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Jitendra",
		Last:  "Kumar",
		Age:   30,
	}
	p2 := person{
		First: "Dharmendra",
		Last:  "Kumar",
		Age:   30,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
