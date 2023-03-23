package main

import "fmt"

func FtoC() {

	fmt.Println("Введите температуру в F:")
	var t float32

	fmt.Scanf("%f", &t)

	output := ((t - 32) * 5 / 9)

	fmt.Println(output)
}
