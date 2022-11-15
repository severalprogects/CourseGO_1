package main

import "fmt"

func main() {
	//Полная запись карты
	var map1 map[int]string = map[int]string{
		1: "Sergey",
		2: "Egor",
	}
	fmt.Println("map1[1]:", map1[1])

	//Сокращенная запись
	map2 := map[int]string{
		1: "Sergey",
		2: "Egor",
	}
	fmt.Println("map2[2]:", map2[2])

	//Проверка на существование элемента
	elem, ok := map2[3]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("No element")
	}
}
