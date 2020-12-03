package main

import(
	"fmt"
)

func main() {
	x := 1
	var y int = 3
	var sum int = x + y

	fmt.Println(sum)

	if sum > 6 {
		fmt.Println("more than 6")
	} else if sum < 6{
		fmt.Println("smaller than 6")
	}
}