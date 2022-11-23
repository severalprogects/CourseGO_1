package main

import "fmt"

//Объявление интерфейса
type talked interface {
	talk(speech string)
}

//Объвления для структур, которые будут реализовывать интерфейс
type people struct {
	name string
}
type robot struct {
	name string
}

//Реализации интерфейса для двух методов
func (p *people) talk(speech string) {
	fmt.Printf("%s talk: %s\n", p.name, speech)
}
func (r *robot) talk(speech string) {
	fmt.Printf("%s talk: %s\n", r.name, speech)
}

//Функция, которая принимает в кач-ве аргумента переменную интерфейса
func hello(t talked) {
	t.talk("Hello")
}

func main() {
	man := &people{name: "Tom"} //Берем указатель на созданный объект структуры,
	hello(man)                  //так как функция, принимающая интерфейс, не может работать с аргументом, переданным по значению

	bot := robot{name: "R2D2"}
	hello(&bot)
}
