package main

import (
	"fmt"
	"sort"
)

func sotrs() {
	fmt.Println(pernumbsgo(33, 44, 6634, 7445, 74, 346, 2, 26, 62))
}

func pernumbsgo(args ...int) int {
	sort.Ints(args)
	a := args[len(args)-1]
	return (a)
}
