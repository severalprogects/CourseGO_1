package main

import "fmt"

func sum(a int, b int) (res int) {
	res = a + b
	return res
}

func allSum(a ...int) (res int) {
	for _, v := range a {
		res = sum(res, v) //Конечно, гораздо проще res += v, но нужно поработать с функциями =)
	}
	return res
}

func print(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func defFunc() {
	num := 5
	defer print(num) //defer принимает то, что перед ней, но выполняется в конце (при чем выполнение отложенных функций происходит из конца к началу)

	//Анонимая функция func(аргументы) {} (то, что мы передаем в качестве аргументов)
	defer func(a int) {
		for i := 0; i < a; i++ {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}(num)

	num = 20
}

func remembVar() func() int {
	myVAr := 2
	myVAr++
	return func() int {
		myVAr++
		return myVAr
	}
}

func main() {
	//Работа с функциями
	num1, num2 := 1, 3
	res := sum(num1, num2)
	print(res)

	//Работа со срезами в функциях
	slc := []int{1, 4, 10}
	print(allSum(slc...)) //чтобы передать слайс, необходимо указать три точки в конце

	//Вызов функции, в которой есть отложенные вызовы
	defFunc()

	//Анонимная функция запоминает свое состояние (мы вернули функцию с инкрементом, она помнит на чем остановилась)
	fu := remembVar()
	for i := 0; i < 3; i++ {
		print(fu())
	}
}
