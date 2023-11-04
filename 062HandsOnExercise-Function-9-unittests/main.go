package main

import "fmt"

func main() {
	fmt.Println(paradise("Hawaii"))
}

func paradise(loc string) string {
	return fmt.Sprint("My Idea of paradise is: ", loc)
}
