package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt" // go get golang.org/x/crypto/bcrypt ---> to get bcrypt
	// to update -> go get -u golang.org/x/crypto/bcrypt
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginpass := `password23`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpass))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN !")
		return
	}
	fmt.Println("YOU ARE LOGGED IN!!")
}
