package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is ", d.first, "and I am walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is ", d.first, "and I am running.")
}

type youngin interface {
	walk()
	run()
}

func youngrun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngrun(d1)

	d2 := &dog{"Padget "}
	d2.walk()
	d2.run()
	youngrun(d2)
}
