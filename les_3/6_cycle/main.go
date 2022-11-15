package main

import "fmt"

func main() {
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
