package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// embedded struct
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "Jitendra",
		last:  "Kaithal",
		age:   30,
	}

	p2 := person{
		first: "Dharmendra",
		last:  "Kaithal",
		age:   30,
	}

	// embeded struct
	sa1 := secretAgent{
		person: person{
			first: "Radhe",
			last:  "Shyam",
			age:   30,
		},
		ltk: true,
	}

	// anonymous struct
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "Ram",
		last:  "Chandra",
		age:   0,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("----------------------------------")
	fmt.Println(sa1)
	fmt.Print(p3)
}
