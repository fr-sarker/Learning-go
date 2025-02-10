package main

import (
	"fmt"
	"sort"
)

func main() {

	b := []int{1, 2, 3, 4, 5, 6}
	a := make([]int, 10)
	copy(a, b)
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 0)
	var i = 2
	copy(b[i+1:], b[i:])
	b[i] = 5
	fmt.Println(b)
	sort.Ints(b)
	fmt.Println(b)
}
