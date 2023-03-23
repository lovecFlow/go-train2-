package main

import "fmt"

func FooToMee() {
	fmt.Println("Введите значение в футах: ")

	var t float32

	fmt.Scanf("%f", &t)

	output := (t * 0.3048)

	fmt.Println(output)
}
