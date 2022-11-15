package main

import "fmt"

func main() {
	//Создадим структуру с двумя полями
	type myStruct struct {
		X int
		Y int
	}

	//Создадим объект на основе структуры и инициализируем его поля
	var myVar1 myStruct = myStruct{X: 1, Y: 10}
	fmt.Println(myVar1)

	//Изменим одно поле для объекта myVar1
	myVar1.X = 11
	fmt.Println(myVar1)

	//Создадим объект, поля которого будут проинициализированны стандартными значениями, затем поменяем одно из них
	myVar2 := myStruct{}
	myVar2.X = 23
	fmt.Println(myVar2)
}
