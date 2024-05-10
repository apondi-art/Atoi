package main

import (
	"fmt"
)

//function PrintInt takes a strint and return an integer

func main() {
	s := "123"
	result := PrintInt(s)
	fmt.Println(result)
}

func PrintInt(s string) int {
	result := 0
	var integer int
	for _, char := range s {
		integer = int(char - '0')

		result = result*10 + integer

	}
	return result
}
