package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

// sorting by age
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// sorting by name
type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := person{"James", 33}
	p2 := person{"Ravi", 32}
	p3 := person{"Rakesh", 13}
	p4 := person{"Bandalu", 45}
	p5 := person{"Mohan", 30}
	p6 := person{"Akash", 19}

	people := []person{p1, p2, p3, p4, p5, p6}

	fmt.Println(people)

	sort.Sort(ByAge(people))

	fmt.Println(people)

	sort.Sort(ByName(people))

	fmt.Println(people)

}
