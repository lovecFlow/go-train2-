package main

import "fmt"

func iter() { //for даёт возможность повторять список инструкций определенное кол-во раз
	i := 1        //Создаётся переменная, в которой хранится число, которое выводится на экран.
	for i <= 10 { //с помощью for создаётся цикл, указывается условное выражение с true или false b
		fmt.Println(i) //И сам блок выполнения
		i = i + 1
	}
}

// Иной вид написания (if)

func incr() {
	for i := 1; i <= 10; i++ { //Вариант с инкрементацией
		if i%2 == 0 { //Оператор % показывает есть ли остаток в выражении
			fmt.Println(i, "even") //Если у i после деления на 2 нет остатка, то "четное"
		} else {
			fmt.Println(i, "odd") //Если у i после деления на 2 есть остаток, то "нечетное"
		}
	}
}

// SWITCH

func switci() {
	switch i { //Начинаем с переключателя (switch) за которым следует выражение (i)
	case 0:
		fmt.Println("Zero") //И серия возможных выражений (case)
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown")
	}
}
