package main

import(
	"fmt"
)

func main() {
	var a []int
	b := []int{1,2,3,4,5}
	a = append(a, 13)
	a[2] = 7
	fmt.Println(a)
	fmt.Println(b)
}