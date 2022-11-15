package main

import "fmt"

func main() {
	//Создадим срез и выведем второй элемент
	var slc1 []int = []int{1, 3, 5, 7}
	fmt.Println("slc1[1]: ", slc1[1])

	//Создадим массив, на основе которого создадим срез и посмотрим его длинну и объем
	arr1 := [...]int{1, 2, 3}
	slc2 := arr1[0:]
	fmt.Printf("slc2 - len: %d, cap: %d\n", len(slc2), cap(slc2))

	//Добавим один элемент к срезу и посмотрим изменение объема и длинны
	slc2 = append(slc2, 4)
	fmt.Printf("slc2 - len: %d, cap: %d\n", len(slc2), cap(slc2))

	//Создадим срез с помощью функции make()
	slc3 := make([]int, 3, 3)
	fmt.Printf("slc3 - len: %d, cap: %d\n", len(slc3), cap(slc3))

	//Добавим элементы в срез из чисел существующего слайса и новых чисел
	slc3 = append(slc2, 5, 6)
	fmt.Println("slc2: ", slc2)
	fmt.Println("slc3: ", slc3)

	//Изменим slc2 и посмотрим изменения для slc2 и slc3
	slc2[1] = 100
	fmt.Println("slc2: ", slc2)
	fmt.Println("slc3: ", slc3)
}
