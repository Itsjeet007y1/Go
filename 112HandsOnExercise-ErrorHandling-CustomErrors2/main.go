package main

import (
	"fmt"
	"log"
)

type sonuMathError struct {
	lat  string
	long string
	err  error
}

func (n sonuMathError) Error() string {
	return fmt.Sprintf("a simple math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.22)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		sme := fmt.Errorf("a simple math error occured")
		return 0, sonuMathError{"5.44", "44.33", sme}
	}
	return 42, nil
}
