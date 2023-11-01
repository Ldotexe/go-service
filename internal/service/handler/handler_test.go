package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"homework-6/internal/service/repository"
	mock_repository "homework-6/internal/service/repository/mocks"
	"homework-6/tests/fixtures"
	"homework-6/tests/states"
)

func Test_Create(t *testing.T) {
	var (
		ctx        context.Context
		student    = fixtures.Student().Build().Ptr()
		studentReq = StudentRequest{
			ID:     student.ID,
			Name:   student.Name,
			Points: student.Points,
		}
	)
	t.Run(
		"success", func(t *testing.T) {
			//arrange
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockStudentRepo(ctrl)

			s := Server{Repo: m}

			m.EXPECT().Add(gomock.Any(), student).Return(nil)
			//act
			status, err := s.Create(ctx, studentReq)
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Run(
				"conflict", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}

					m.EXPECT().Add(gomock.Any(), student).Return(repository.ErrConflict)
					//act
					status, err := s.Create(ctx, studentReq)
					//assert
					require.Equal(t, http.StatusConflict, status)
					require.ErrorIs(
						t, err, repository.ErrConflict,
					)
				},
			)
			t.Run(
				"internal error", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}

					m.EXPECT().Add(gomock.Any(), student).Return(assert.AnError)
					//act
					status, err := s.Create(ctx, studentReq)
					//assert
					require.Equal(t, http.StatusInternalServerError, status)
					require.ErrorIs(t, err, assert.AnError)
				},
			)
		},
	)
}

func Test_Get(t *testing.T) {
	var (
		ctx context.Context
		id  = states.StudentID1
	)
	t.Run(
		"success", func(t *testing.T) {
			//arrange
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockStudentRepo(ctrl)

			s := Server{Repo: m}
			m.EXPECT().GetByID(gomock.Any(), int64(id)).Return(fixtures.Student().Build().Ptr(), nil)

			//act
			answer, status, err := s.Get(ctx, int64(id))
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			require.JSONEq(t, `{"ID":28,"Name":"Aleksei","Points":60}`, string(answer))
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"not found", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}
					m.EXPECT().GetByID(gomock.Any(), int64(id)).Return(nil, repository.ErrObjectNotFound)

					//act
					answer, status, err := s.Get(ctx, int64(id))
					//assert
					require.Equal(t, http.StatusNotFound, status)
					require.ErrorIs(t, err, repository.ErrObjectNotFound)
					require.Nil(t, answer)
				},
			)
			t.Run(
				"internal error", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}
					m.EXPECT().GetByID(gomock.Any(), int64(id)).Return(nil, assert.AnError)

					//act
					answer, status, err := s.Get(ctx, int64(id))
					//assert
					require.Equal(t, http.StatusInternalServerError, status)
					require.ErrorIs(t, err, assert.AnError)
					require.Nil(t, answer)
				},
			)
		},
	)
}

func Test_Update(t *testing.T) {
	var (
		ctx        context.Context
		student    = fixtures.Student().Build().Ptr()
		studentReq = StudentRequest{
			ID:     student.ID,
			Name:   student.Name,
			Points: student.Points,
		}
	)
	t.Run(
		"success", func(t *testing.T) {
			//arrange
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockStudentRepo(ctrl)

			s := Server{Repo: m}

			m.EXPECT().Update(gomock.Any(), student).Return(nil)
			//act
			status, err := s.Update(ctx, studentReq)
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"not found", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}

					m.EXPECT().Update(gomock.Any(), student).Return(repository.ErrObjectNotFound)
					//act
					status, err := s.Update(ctx, studentReq)
					//assert
					require.Equal(t, http.StatusNotFound, status)
					require.ErrorIs(t, err, repository.ErrObjectNotFound)
				},
			)
			t.Run(
				"internal error", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}

					m.EXPECT().Update(gomock.Any(), student).Return(assert.AnError)
					//act
					status, err := s.Update(ctx, studentReq)
					//assert
					require.Equal(t, http.StatusInternalServerError, status)
					require.ErrorIs(t, err, assert.AnError)
				},
			)
		},
	)
}

func Test_Delete(t *testing.T) {
	var (
		ctx context.Context
		id  = states.StudentID1
	)
	t.Run(
		"success", func(t *testing.T) {
			//arrange
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockStudentRepo(ctrl)

			s := Server{Repo: m}

			m.EXPECT().Delete(gomock.Any(), int64(id)).Return(nil)
			//act
			status, err := s.Delete(ctx, int64(id))
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"not found", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}

					m.EXPECT().Delete(gomock.Any(), int64(id)).Return(repository.ErrObjectNotFound)
					//act
					status, err := s.Delete(ctx, int64(id))
					//assert
					require.Equal(t, http.StatusNotFound, status)
					require.ErrorIs(t, err, repository.ErrObjectNotFound)
				},
			)
			t.Run(
				"internal error", func(t *testing.T) {
					//arrange
					t.Parallel()
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()
					m := mock_repository.NewMockStudentRepo(ctrl)

					s := Server{Repo: m}

					m.EXPECT().Delete(gomock.Any(), int64(id)).Return(assert.AnError)
					//act
					status, err := s.Delete(ctx, int64(id))
					//assert
					require.Equal(t, http.StatusInternalServerError, status)
					require.ErrorIs(t, err, assert.AnError)
				},
			)
		},
	)
}
