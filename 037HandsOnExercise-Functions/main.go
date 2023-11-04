package main

import "fmt"

func main() {
	foo()
	bar("Jitendra")
	s := aloha("Jitu")
	fmt.Println(s)

	s1, n := twoParamsAndReturn("Todd", 40)
	fmt.Println(s1, n)
}

// no params no return
func foo() {
	fmt.Println("I am from func foo.")
}

// 1 param no return
func bar(s string) {
	fmt.Println("My name is:", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me mr.", s)
}

func twoParamsAndReturn(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "Todd is this old in dog years"), age
}
