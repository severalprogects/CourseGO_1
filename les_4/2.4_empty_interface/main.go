package main

import "fmt"

type myStruct struct {
	myLine string
}

//Функция принимает пустой интерфейс, следовательно может работать с любым типом и с любой структурой (все это реализует пустой интерфейс)
//Но это не весь функционал пустых интерфейсов, дополнение добавлю позднее
func myPrint(i interface{}) {
	switch i.(type) {
	case int:
		i = i.(int)
		fmt.Println("Enter int number is:", i)
	case float64:
		f := i.(float64)
		fmt.Println("Enter float65 number is:", f)
	case string:
		s := i.(string)
		fmt.Println("Enter string line is:", s)
	case myStruct:
		m := i.(myStruct)
		fmt.Println("Entr myStruct line is:", m.myLine)
	}
}

func main() {
	var i int = 100
	var f float64 = 10.11
	var s string = "string"
	var m myStruct = myStruct{myLine: "myString"}

	myPrint(i)
	myPrint(f)
	myPrint(s)
	myPrint(m)
}
