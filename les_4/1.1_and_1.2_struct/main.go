package main

import "fmt"

type people struct {
	name string
	age  int
}

type gender struct {
	sex string
}

type student struct {
	idStud  int
	infStud people
	infSex  *gender
}

func updateName(s student) {
	s.infStud.name = "Tom"
}

func updateNameWithPtr(s *student) {
	s.infStud.name = "Tom"
}

func changeSex(s student) {
	if s.infSex.sex == "male" {
		s.infSex.sex = "female"
	} else {
		s.infSex.sex = "male"
	}
}

func main() {
	stud := student{}
	fmt.Println(stud) //все значения будут приравнены к "нулевым" для каждого типа

	stud.infStud.name = "Jack"
	stud.infStud.age = 18
	stud.idStud = 1614
	fmt.Println(stud)

	stud.infSex = new(gender) //необходимо выделить память, т.к. изначально nil (или ... = &gender{})
	stud.infSex.sex = "male"
	fmt.Println(stud)            //в конце указан адрес области памяти, где записан пол
	fmt.Println(stud.infSex.sex) //чтобы посмотреть, что записано по адресу не нужно выполнять операцию разыменования

	updateName(stud) //Пытаемся обновить значение
	fmt.Println(stud)

	updateNameWithPtr(&stud) //Обновляем значение, используя указатель
	fmt.Println(stud)

	//Изменяем пол (передача аргумента по значению, однако мы меняем поле - это связанно с тем,
	//что скопированный аргумент имеет тот же указатель на область памяти, где содердится информация о поле (sex) оригинального объекта)
	changeSex(stud)
	fmt.Println(stud.infSex.sex)
}
