package student

type Student struct {
	idStud int
	name   string
}

func (s *Student) SetName(name string) {
	s.name = name
}

func NewStudent(name string) Student {
	return Student{
		idStud: 7,
		name:   name,
	}
}
