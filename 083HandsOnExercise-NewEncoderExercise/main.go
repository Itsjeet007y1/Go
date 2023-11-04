package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	j := `[{"First":"Alok","Last":"Marya","Age":11},{"First":"Vipin","Last":"Kumar","Age":11},{"First":"Gaurav","Last":"Gupta","Age":11}]`

	bs := []byte(j)
	student := []student{}

	err := json.Unmarshal(bs, &student)
	if err != nil {
		fmt.Println(err)
	}

	for _, s := range student {
		fmt.Println(s.First, s.Last, s.Age)
	}

	err = json.NewEncoder(os.Stdout).Encode(student)

	if err != nil {
		fmt.Println("Something went wrong...")
	}

	fmt.Println(student)
}
