package main

import "fmt"

const PI = 3.14

type sideRectangle float64 //Сторона прямоугольника
type radius float64        //Радиус круга

type figure interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	name string
	a, b sideRectangle
}
type Circle struct {
	name string
	r    radius
}

func (r *Rectangle) Area() float64 {
	return float64(r.a) * float64(r.b)
}
func (r *Rectangle) Type() string {
	return r.name
}

func (c *Circle) Area() float64 {
	return PI * float64(c.r) * float64(c.r)
}
func (c *Circle) Type() string {
	return c.name
}

func main() {
	myRect := Rectangle{name: "Rectangle", a: 5, b: 6}
	fmt.Printf("Figure %s has area: %.2f\n", myRect.Type(), myRect.Area())

	myCirc := Circle{name: "Circle", r: 10}
	fmt.Printf("Figure %s has area: %.2f\n", myCirc.Type(), myCirc.Area())
}
