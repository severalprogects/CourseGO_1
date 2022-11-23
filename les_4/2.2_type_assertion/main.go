package main

import "fmt"

type vehicle interface {
	open(key string)
	ride() string
}

type car struct {
	key    string
	isOpen bool
	color  string
}
type bus struct {
	key    string
	isOpen bool
	color  string
}

func (c *car) open(key string) {
	if c.key == key {
		c.isOpen = true
	}
}
func (c *car) ride() string {
	if c.isOpen {
		return "Car is ride"
	}
	return "Car is not open"
}

func (b *bus) open(key string) {
	if b.key == key {
		b.isOpen = true
	}
}
func (b *bus) ride() string {
	if b.isOpen {
		return "Bus is ride"
	}
	return "Bus is not open"
}

func script(v vehicle, key string) {
	// Скрипт: открываем машину ключем и пытаемся на ней поехать
	v.open(key)
	fmt.Println(v.ride())

	// Вариант утверждения типа с помощью возвращения двух аргументов и проверке в блоке if
	cars, ok := v.(*car)
	if ok {
		fmt.Println(cars.color)
	}

	// Вариант утверждения типа через оператор switch
	switch v.(type) {
	case *car:
		cars = v.(*car)
		fmt.Println("Color cur is", cars.color)
	case *bus:
		buses := v.(*bus)
		fmt.Println("Color bus is", buses.color)
	}
}

func main() {
	ladaCar := car{key: "carKey", isOpen: false, color: "red"}
	script(&ladaCar, "carKey")

	volvoBus := bus{key: "busKey", isOpen: false, color: "green"}
	script(&volvoBus, "carKey")
}
