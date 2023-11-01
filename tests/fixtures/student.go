package fixtures

import (
	"homework-6/internal/service/repository"
	"homework-6/tests/states"
)

type StudentBuilder struct {
	instance *repository.Student
}

func Student() *StudentBuilder {
	return &StudentBuilder{instance: &repository.Student{}}
}

func (s *StudentBuilder) ID(v int64) *StudentBuilder {
	s.instance.ID = v
	return s
}

func (s *StudentBuilder) Name(v string) *StudentBuilder {
	s.instance.Name = v
	return s
}

func (s *StudentBuilder) Points(v int64) *StudentBuilder {
	s.instance.Points = v
	return s
}

func (s *StudentBuilder) Ptr() *repository.Student {
	return s.instance
}

func (s *StudentBuilder) Value() repository.Student {
	return *s.instance
}

func (s *StudentBuilder) Build() *StudentBuilder {
	return Student().ID(states.StudentID1).Name(states.StudentName1).Points(states.StudentPoints1)
}
