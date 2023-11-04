package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeout(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {

	p := person{
		first: "Jenny",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeout(f)
	p.writeout(&b)
	fmt.Println(b.String())

}
