//go:build integration
// +build integration

package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"homework-6/internal/service/repository"
	"homework-6/internal/service/repository/studentDB"
	"homework-6/tests/fixtures"
	"homework-6/tests/states"
)

func TestAddStudent(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			var ctx = context.Background()
			db.SetUp(t)
			defer db.TearDown()
			// arrange

			repo := studentDB.NewStudents(db.DB)
			//act
			err := repo.Add(ctx, fixtures.Student().Build().Ptr())
			//assert
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
					//act
					err := repo.Add(ctx, fixtures.Student().Build().Ptr())
					//assert
					require.NoError(t, err)
					//act
					err = repo.Add(ctx, fixtures.Student().Build().Ptr())
					//assert
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
			//act
			err := repo.Add(ctx, fixtures.Student().Build().Ptr())
			//assert
			require.NoError(t, err)
			//act
			res, err := repo.GetByID(ctx, states.StudentID1)
			//assert
			require.NoError(t, err)
			require.Equal(t, fixtures.Student().Build().Ptr(), res)
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
					//act
					student, err := repo.GetByID(ctx, states.StudentID1)
					//assert
					require.ErrorIs(
						t, err, repository.ErrObjectNotFound,
					)
					require.Nil(t, student)
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
			//act
			err := repo.Add(ctx, fixtures.Student().Build().Ptr())
			//assert
			require.NoError(t, err)
			//act
			err = repo.Update(
				ctx,
				fixtures.Student().ID(states.StudentID1).Name(states.StudentName2).Points(states.StudentPoints2).Ptr(),
			)
			//assert
			require.NoError(t, err)
			//act
			res, err := repo.GetByID(ctx, states.StudentID1)
			//assert
			require.NoError(t, err)
			require.Equal(
				t,
				fixtures.Student().ID(states.StudentID1).Name(states.StudentName2).Points(states.StudentPoints2).Ptr(),
				res,
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
					//act
					err := repo.Update(
						ctx,
						fixtures.Student().ID(states.StudentID1).Name(states.StudentName2).Points(states.StudentPoints2).Ptr(),
					)
					//assert
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
			//act
			err := repo.Add(ctx, fixtures.Student().Build().Ptr())
			//assert
			require.NoError(t, err)
			//act
			err = repo.Delete(ctx, states.StudentID1)
			//assert
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
					//act
					err := repo.Delete(ctx, states.StudentID1)
					//assert
					require.ErrorIs(t, err, repository.ErrObjectNotFound)
				},
			)
		},
	)
}
