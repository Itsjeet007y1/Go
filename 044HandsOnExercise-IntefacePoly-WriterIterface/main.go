package main

import (
	"bytes"
	"fmt"
)

// type person struct {
// 	first string
// }

// func (p person) writeout(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()

	// s := []byte("Hello gophers !")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	// Wrting to a byte buffer

	b := bytes.NewBufferString("Hello ")

	fmt.Println(b.String())

	b.WriteString("Hello Gophers !")

	fmt.Println(b.String())

	b.Reset()

	b.WriteString("Buffer has been reset.")

	fmt.Println(b.String())

	b.Reset()

	b.Write([]byte("Happy Happy"))

	fmt.Println(b.String())

}
