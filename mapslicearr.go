package main

import (
	"fmt"
	"sort"
)

func abcdef() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	z := x[2:5]
	fmt.Println(z)

	fmt.Println(x[4])

	e := make([]int, 3, 9)

	fmt.Println(len(e))

	j := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	sort.Ints(j)
	fmt.Println(j[len(j)-1])

}
