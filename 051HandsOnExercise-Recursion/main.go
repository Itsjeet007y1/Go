package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(factorial(0))
	fmt.Println(reverseString("Jitendra"))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func reverseString(s string) string {
	str := ""
	str1 := strings.Split(s, "")
	fmt.Println(len(s))
	fmt.Println(str1[7])
	if len(s) == 0 {
		return ""
	}
	for i := len(s) - 1; i >= 0; i-- {
		str = str + str1[i]
	}
	return str
}
