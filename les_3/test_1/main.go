/*
 * Задание:
 * Создать две константы со статусами проверки.
 * Создать структуру, которая содержит ID и статус проверки.
 * Создать ф-ию, которая возвращает слайс структур, которые были предварительно инициализированны ID и стаутсом.
 * Проверить принимаемй слайс структур, и вывести только те ID, для которых стаутс равен "pass".
 */

package main

import "fmt"

const PassStatus = "pass"
const FailStatus = "fail"

type heatleCheck struct {
	serviceId int
	status    string
}

func generateCheck() (slice []heatleCheck) {
	for i := 0; i < 5; i++ {
		var check heatleCheck
		if i%2 == 0 {
			check.serviceId = i
			check.status = PassStatus
		} else {
			check.serviceId = i
			check.status = FailStatus
		}
		slice = append(slice, check)
	}
	return slice
}

func main() {
	slice := generateCheck()
	for _, v := range slice {
		if v.status == PassStatus {
			fmt.Println(v.serviceId)
		}
	}
}
