package main

import "fmt"

//Карта или ассоциативный массив или словарь -- неупорядочненая коллекция пар вида ключ-значение.

func stringmap() {
	// var x map[string]int //Карта стрингов для int-ов
	// //Карта представляется в связке с клбчевым словом map, следуюшим за ним типом ключа в скобках  и типом значения после скобок.
	// x["key"] = 10
	// fmt.Println(x)

	// //Результатом будет ошибка, тк карту необходимо инициализировтаь. Написать надо вот так:

	y := make(map[string]int)
	y["key"] = 10
	fmt.Println(y["key"])

	//Результатом будет 10. Похоже на массивы, но ключом тут является не число, а строка, тк в карте указаг тип string.

	z := make(map[int]int)
	z[1] = 10    //В самом начале длина 0, но  после создания z[1] = 10 она станет равна 1.
	delete(z, 1) //Можем удалить элементы из карты используя delete

	fmt.Println(z[1])   //Длина карты может меняться, когда в неё добавляется новый элемент.
	fmt.Println(len(z)) //В картах отсчёт идёт с 1, тк не является последовательностью

	//ПРИМЕР
	//elements -- мапа с 10 химическими элементами, индексируемых симовлами. Мапы зачастую используют как словари или таблицы.
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

}
