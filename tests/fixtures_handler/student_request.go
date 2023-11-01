package fixtures_handler

import (
	"homework-6/internal/service/handler"
	"homework-6/tests/states"
)

type StudentReqBuilder struct {
	instance *handler.StudentRequest
}

func StudentReq() *StudentReqBuilder {
	return &StudentReqBuilder{instance: &handler.StudentRequest{}}
}

func (s *StudentReqBuilder) ID(v int64) *StudentReqBuilder {
	s.instance.ID = v
	return s
}

func (s *StudentReqBuilder) Name(v string) *StudentReqBuilder {
	s.instance.Name = v
	return s
}

func (s *StudentReqBuilder) Points(v int64) *StudentReqBuilder {
	s.instance.Points = v
	return s
}

func (s *StudentReqBuilder) Ptr() *handler.StudentRequest {
	return s.instance
}

func (s *StudentReqBuilder) Value() handler.StudentRequest {
	return *s.instance
}

func (s *StudentReqBuilder) Build() *StudentReqBuilder {
	return StudentReq().ID(states.StudentID1).Name(states.StudentName1).Points(states.StudentPoints1)
}
