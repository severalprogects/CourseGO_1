package main

import "fmt"

type sender interface {
	send(email string)
}
type caller interface {
	call(number string)
}

// Встраиваем интерфейсы для отправки сообщения и звонка в новый интерфейс
type phone interface {
	sender
	caller
	play(music int) // Мы можем добавлять дополнительное поведение для интерфейсов
}

// Реализуем поведение встроенных интерфейсов, а также метод самого интерфейса
type iPhone struct {
	name  string
	color string
}

func (i *iPhone) send(email string) {
	fmt.Printf("Email: \"%s\" - delivered\n", email)
}
func (i *iPhone) call(number string) {
	fmt.Printf("Call by number: \"%s\" - is made\n", number)
}
func (i *iPhone) play(music int) {
	switch music {
	case 1:
		fmt.Println("Music: \"Bara-bara-bara\"")
	case 2:
		fmt.Println("Music: \"Bere-bere-bere\"")
	default:
		fmt.Println("Music not found")
	}
}

func main() {
	iPhone11 := iPhone{name: "iPhone11", color: "Red"}
	iPhone11.call("+78009303030")
	iPhone11.send("Hello, how are you?")
	iPhone11.play(1)
}
