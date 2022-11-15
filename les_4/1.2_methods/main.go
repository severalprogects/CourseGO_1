package main

import "fmt"

type myStruct struct {
	name string
	age  int
}

func (m myStruct) newName(name string) {
	m.name = name
}

func (m *myStruct) newAge(age int) {
	m.age = age
}

func main() {
	myVar := &myStruct{}
	myVar.newName("Tom") //Не можем изменить, так как передаем аргумент по значению
	fmt.Println(myVar)

	myVar.name = "Jack"
	myVar.newAge(44) //Изменяем, так как передаем по указателю
	fmt.Println(myVar)
}
