package student

type person struct {
	name string
	age  uint
}

type Student struct {
	person
	studID uint
}

func (s *Student) Name() string {
	return s.name
}

func (s *Student) Age() uint {
	return s.age
}

func (s *Student) StudID() uint {
	return s.studID
}

func NewStudent(name string, age uint, studID uint) Student {
	return Student{
		person: person{
			name: name,
			age:  age,
		},
		studID: studID,
	}
}
