package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error happened..", err)
		log.Println("error happened..", err)
		log.Fatalln(err)
		log.Panicln(err)
		panic(err)
	}
}

// package log implements a simple loggin package.. writes to standard error and
// and prints the date and time of each loggin messages

// log.println calls output to print to the standartd logger. Arguments are handelled in
// the manner of fmt.Println.
