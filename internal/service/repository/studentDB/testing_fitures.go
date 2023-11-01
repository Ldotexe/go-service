package studentDB

import (
	"testing"

	"go.uber.org/mock/gomock"
	mock_db "homework-6/internal/service/db/mocks"
)

type studentRepoFixtures struct {
	ctrl   *gomock.Controller
	mockDB mock_db.MockDBops
	repo   StudentsRepo
}

func setUP(t *testing.T) studentRepoFixtures {
	ctrl := gomock.NewController(t)
	mockDB := mock_db.NewMockDBops(ctrl)
	repo := NewStudents(mockDB)
	return studentRepoFixtures{
		ctrl:   ctrl,
		mockDB: *mockDB,
		repo:   *repo,
	}
}

func (s *studentRepoFixtures) tearDown() {
	s.ctrl.Finish()
}
