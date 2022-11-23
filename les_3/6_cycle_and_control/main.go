package main

import "fmt"

func main() {
	//Оператор if
	number := 3
	if number < 0 {
		fmt.Println("Выражение в блоке if меньше 0")
	} else if number == 1 || number == 2 {
		fmt.Println("Выражение в блоке if равно 1 или 2")
	} else {
		fmt.Println("Выражение в блоке if больше 2")
	}

	//Оператор switch (в данном операторе мы не проваливаемся на уровень ниже, если не запишем оператор break)
	switch number {
	case 1:
		fmt.Println("Выражение в блоке switch равно 1")
	case 3:
		fmt.Println("Выражение в блоке switch равно 3")
	default:
		fmt.Println("Значение не опознано")
	}

	//Еще один вариант использования switch
	switch {
	case number < 0:
		fmt.Println("Выражение в блоке if меньше 0")
	case number == 1 || number == 2:
		fmt.Println("Выражение в блоке if равно 1 или 2")
	default:
		fmt.Println("Выражение в блоке if больше 2")
	}

	//Бесконечный цикл (если бы не было условия для выхода)
	i := 0
	for {
		i++
		fmt.Print(i)

		if i == 5 {
			break
		}
	}
	fmt.Println()

	//Цикл с условием
	j := 0
	for j < 5 {
		j++
		fmt.Print(j)
	}
	fmt.Println()

	//Цикл с параметром
	for k := 0; k < 3; k++ {
		fmt.Print(k)
	}
	fmt.Println()
	arr := [...]int{1, 4, 8}

	//Цикл range
	for _, v := range arr {
		fmt.Print(v)
	}

}
