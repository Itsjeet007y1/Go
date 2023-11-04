package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	First string
	Last  string
	Age   int
}

func main() {
	s1 := student{
		First: "Alok",
		Last:  "Marya",
		Age:   11,
	}
	s2 := student{
		First: "Vipin",
		Last:  "Kumar",
		Age:   11,
	}
	s3 := student{
		First: "Gaurav",
		Last:  "Gupta",
		Age:   11,
	}

	students := []student{s1, s2, s3}

	fmt.Println(students)

	s, err := json.Marshal(students)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(s))
}
