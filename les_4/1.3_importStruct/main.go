package main

import (
	"fmt"

	"github.com/severalprogects/CourseGO_1/les_4/1.3_importStruct/student"
)

func main() {
	myStud := student.NewStudent("Tom")
	fmt.Println(myStud)
}
