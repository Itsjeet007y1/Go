package main

import "fmt"

func main() {

	m := make(map[string][]string)

	m[`jitendra`] = []string{`this`, `is`, `my`, `first`, `name`}
	m[`kumar`] = []string{`this`, `is`, `my`, `middle`, `name`}
	m[`kaithal`] = []string{`this`, `is`, `my`, `last`, `name`}
	m[`pintu`] = []string{`this`, `is`, `my`, `another`, `name`}

	fmt.Printf("%#v\n", m)

	fmt.Println("-----------------------------------------------")

	for k, v := range m {
		fmt.Println(k)
		for i, e := range v {
			fmt.Println(i, " ", e)
		}
	}

	// deleting a record from map
	delete(m, `pintu`)

	fmt.Println("---------------------record deleted---------------------")

	for k, v := range m {
		fmt.Println(k)
		for i, e := range v {
			fmt.Println(i, " ", e)
		}
	}
}
