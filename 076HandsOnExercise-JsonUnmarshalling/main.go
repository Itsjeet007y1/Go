package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Jitendra","Last":"Kumar","Age":30},{"First":"Dharmendra","Last":"Kumar","Age":30}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people := []person{}
	// var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for _, v := range people {
		fmt.Println(v.First, v.Last, v.Age)
	}
}
