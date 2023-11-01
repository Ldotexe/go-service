//go:build integration_handler
// +build integration_handler

package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"homework-6/internal/service/handler"
	"homework-6/internal/service/repository"
	"homework-6/internal/service/repository/studentDB"
	"homework-6/tests/fixtures_handler"
	"homework-6/tests/states"
)

func TestCreate(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			var ctx = context.Background()
			db.SetUp(t)
			defer db.TearDown()
			// arrange
			repo := studentDB.NewStudents(db.DB)
			s := handler.Server{Repo: repo}
			//act
			status, err := s.Create(ctx, fixtures_handler.StudentReq().Build().Value())
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"conflict", func(t *testing.T) {
					var ctx = context.Background()
					db.SetUp(t)
					defer db.TearDown()
					//arrange
					repo := studentDB.NewStudents(db.DB)
					s := handler.Server{Repo: repo}
					//act
					status, err := s.Create(ctx, fixtures_handler.StudentReq().Build().Value())
					//assert
					require.Equal(t, http.StatusOK, status)
					require.NoError(t, err)
					//act
					status, err = s.Create(ctx, fixtures_handler.StudentReq().Build().Value())
					//assert
					require.Equal(t, http.StatusConflict, status)
					require.ErrorIs(
						t, err, repository.ErrConflict,
					)
				},
			)
		},
	)
}

func TestGetStudent(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			var ctx = context.Background()
			db.SetUp(t)
			defer db.TearDown()
			// arrange
			repo := studentDB.NewStudents(db.DB)
			s := handler.Server{Repo: repo}
			//act
			status, err := s.Create(ctx, fixtures_handler.StudentReq().Build().Value())
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			//act
			res, status, err := s.Get(ctx, states.StudentID1)
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			require.JSONEq(t, "{\"ID\":28,\"Name\":\"Aleksei\",\"Points\":60}", string(res))
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"not found", func(t *testing.T) {
					var ctx = context.Background()
					db.SetUp(t)
					defer db.TearDown()
					//arrange
					repo := studentDB.NewStudents(db.DB)
					s := handler.Server{Repo: repo}
					//act
					res, status, err := s.Get(ctx, states.StudentID1)
					//assert
					require.Equal(t, http.StatusNotFound, status)
					require.ErrorIs(
						t, err, repository.ErrObjectNotFound,
					)
					require.Nil(t, res)
				},
			)
		},
	)
}

func TestUpdateStudent(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			var ctx = context.Background()
			db.SetUp(t)
			defer db.TearDown()
			// arrange
			repo := studentDB.NewStudents(db.DB)
			s := handler.Server{Repo: repo}
			//act
			status, err := s.Create(ctx, fixtures_handler.StudentReq().Build().Value())
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			//act
			status, err = s.Update(
				ctx,
				fixtures_handler.StudentReq().ID(states.StudentID1).Name(states.StudentName2).Points(states.StudentPoints2).Value(),
			)
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			//act
			res, status, err := s.Get(ctx, states.StudentID1)
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			require.JSONEq(
				t, "{\"ID\":28,\"Name\":\"Alex\",\"Points\":30}",
				string(res),
			)
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"not found", func(t *testing.T) {
					var ctx = context.Background()
					db.SetUp(t)
					defer db.TearDown()
					//arrange
					repo := studentDB.NewStudents(db.DB)
					s := handler.Server{Repo: repo}
					//act
					status, err := s.Update(
						ctx,
						fixtures_handler.StudentReq().ID(states.StudentID1).Name(states.StudentName2).Points(states.StudentPoints2).Value(),
					)
					//assert
					require.Equal(t, http.StatusNotFound, status)
					require.ErrorIs(t, err, repository.ErrObjectNotFound)
				},
			)
		},
	)
}
func TestDeleteStudent(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			var ctx = context.Background()
			db.SetUp(t)
			defer db.TearDown()
			// arrange
			repo := studentDB.NewStudents(db.DB)
			s := handler.Server{Repo: repo}
			//act
			status, err := s.Create(ctx, fixtures_handler.StudentReq().Build().Value())
			//assert
			require.Equal(t, http.StatusOK, status)
			require.NoError(t, err)
			//act
			status, err = s.Delete(ctx, states.StudentID1)
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
					var ctx = context.Background()
					db.SetUp(t)
					defer db.TearDown()
					//arrange
					repo := studentDB.NewStudents(db.DB)
					s := handler.Server{Repo: repo}
					//act
					status, err := s.Delete(ctx, states.StudentID1)
					//assert
					require.Equal(t, http.StatusNotFound, status)
					require.ErrorIs(
						t, err, repository.ErrObjectNotFound,
					)
				},
			)
		},
	)
}
