package main

import (
	"errors"
	"fmt"
)

type client struct {
	name string
}

type server struct {
	serverOn bool
}
type serverError struct {
	err error
}

// NewClient() - функция-конструктор, для создания нового клиента
func NewClient(name string, serv server) (client, error) {
	//Если сервер выключен
	if !serv.serverOn {
		//Создание ошибки с помощью созданной функции
		err := newServerError("Ошибка подключения к серверу")
		return client{}, err
	}
	//Если перменная, содеражщая имя, пустая
	if name == "" {
		err := errors.New("Отсутствие имени клиента")
		return client{}, err
	}
	return client{name: name}, nil
}

// CopyClient - функция копирования, для создания нового клиента на основе данных об уже существующем
func CopyClient(oldClient client, serv server) (client, error) {
	newClient, err := NewClient(oldClient.name, serv)
	if err != nil {
		//Создание вложенной ошибки с помощью функции Errorf() из пакета fmt
		err = fmt.Errorf("Ошибка копирования информации о клиенте: %w", err)
		return newClient, err
	}
	return newClient, nil
}

// NewServer() - функция-конструктор для создания нового сервера
func NewServer(servOn bool) server {
	return server{serverOn: servOn}
}

// change() - функция, которая изменяет состояние срвера
func (s *server) change() {
	s.serverOn = !s.serverOn
}

// Описание метода Error() для реализации интерфеса error структурой serverError
func (s *serverError) Error() string {
	return s.err.Error()
}

// newServerError - функция-конструктор, для создания собственной ошибки
func newServerError(msg string) error {
	return &serverError{err: errors.New(msg)}
}

func main() {
	//Создание своего "сервера"
	myServer := NewServer(false)
	//Создание слайса, для хранения всех клиентов ("база данных" о всех клиентах)
	persons := []client{}

	//Проверяем полученное значение ошибки, и если оно не равно нулю, то выводим ее значение,
	//иначе добавляем полученного клиента в "бд"
	person, err := NewClient("Jack", myServer)
	if err != nil {
		fmt.Println(err)
	} else {
		persons = append(persons, person)
	}

	fmt.Println("-----") // Разделитель вывода "----------"

	//Включаем "сервер"
	myServer.change()
	//Передаем пустое значение для аргумента, содержащего имя
	person, err = NewClient("", myServer)
	if err != nil {
		fmt.Println(err)
	} else {
		persons = append(persons, person)
	}

	fmt.Println("-----") // Разделитель вывода "----------"

	person, err = NewClient("Tom", myServer)
	if err != nil {
		fmt.Println(err)
	} else {
		persons = append(persons, person)
	}
	//Выводим значение, записанное "БД"
	fmt.Println("Значение первого клиента в БД:", persons[0].name)

	fmt.Println("-----") // Разделитель вывода "----------"

	//Выключаем "сервер"
	myServer.change()
	//Пытаемся скопировать клиента
	person, err = CopyClient(persons[0], myServer)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err)) //Каждый последующий вызов ф-ции Unwrap() все глубже заходит во вложенную ошибку
		if errors.As(err, new(*serverError)) {
			fmt.Println("Ошибки совпадают")
		}
		myError := newServerError("Ошибка подключения к серверу") //Создадим ошибку типа serverError
		if errors.Is(err, myError) {
			fmt.Println("Ошибка типа: serverError") //Ошибки не совпадают - вывода не будет (errors.Unwrap(err) != myError)
		}
	} else {
		persons = append(persons, person)
	}

	fmt.Println("-----") // Разделитель вывода "----------"

	//Включаем "сервер"
	myServer.change()
	//Обнуляем значение имени
	persons[0].name = ""
	//Пытаемся скопировать клиента
	person, err = CopyClient(persons[0], myServer)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
		if errors.As(err, new(*serverError)) {
			fmt.Println("Ошибки совпадают") //В данном случае ошибки совпадать не будут
		}
	} else {
		persons = append(persons, person)
	}
}
