package main

import (
	"fmt"
)

func half(n int) (int, bool) {
	halfn := n / 2                             //Обозначем деление
	return halfn, halfn*2 == n && halfn%2 == 0 //Возвращаем значения (указания выше)
}

func println() { //Выносим в мейн.
	fmt.Println(half(20))
	fmt.Println(half(10))
	fmt.Println(half(3))
	fmt.Println(half(2))
}
