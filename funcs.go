package main

import "fmt"

//Глобальные и локальные переменные. Тут просто: локальные перемнные работают в области функции (func main), а глобальные работают в области всего пакета (package). Также можно переинициализировать их путём поставления =.

var a, b, c int

func abc() {
	a, b, c := 1, 2, 3 //Если убрать : то значения уйдут в верхнюю инициализацию, а если оставить, то значения будут работать локально в функции main. В var и func main a, b, c имеют разные значения в данный момент.

	fmt.Println(a, b, c)
}

//Инкапсуляция

func args() {
	print("Вызов 1") //Мы не можем вызвать функции без аргумента.
	print("Вызов 1")
	print("Вызов 1")
}

func print(message string) { //		Инкапсулировали 			     использовали свреху

	fmt.Println("func")
}

//Возвращаемые значения

func sayName(name string) string {
	return "Hello" + name
}
