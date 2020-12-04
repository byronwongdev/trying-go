package main

import (
	"fmt"
)

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// using range

	arr := []string{"a", "b", "c", "d", "e"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
}
