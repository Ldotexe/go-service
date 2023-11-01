// For training

package studentDB

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"homework-6/internal/service/repository"
	"homework-6/tests/states"
)

func TestStudentsRepo_GetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = states.StudentID1
	)
	t.Run(
		"success", func(t *testing.T) {
			t.Parallel()
			//arrange
			s := setUP(t)
			defer s.tearDown()

			s.mockDB.EXPECT().Get(
				gomock.Any(), gomock.Any(), "SELECT id,name,points FROM students WHERE id=$1", gomock.Any(),
			).Return(nil)
			//act
			user, err := s.repo.GetByID(ctx, int64(id))
			//assert

			require.NoError(t, err)
			assert.Equal(t, int64(0), user.ID)
		},
	)
	t.Run(
		"fail", func(t *testing.T) {
			t.Parallel()
			t.Run(
				"fail, not found", func(t *testing.T) {
					t.Parallel()
					//arrange
					s := setUP(t)
					defer s.tearDown()

					s.mockDB.EXPECT().Get(
						gomock.Any(), gomock.Any(), "SELECT id,name,points FROM students WHERE id=$1", gomock.Any(),
					).Return(pgx.ErrNoRows)
					//act
					user, err := s.repo.GetByID(ctx, int64(id))
					//assert
					require.ErrorIs(t, err, repository.ErrObjectNotFound)
					assert.Nil(t, user)
				},
			)
			t.Run(
				"fail, internal error", func(t *testing.T) {
					t.Parallel()
					//arrange
					s := setUP(t)
					defer s.tearDown()

					s.mockDB.EXPECT().Get(
						gomock.Any(), gomock.Any(), "SELECT id,name,points FROM students WHERE id=$1", gomock.Any(),
					).Return(assert.AnError)
					//act
					user, err := s.repo.GetByID(ctx, int64(id))
					//assert
					require.ErrorIs(t, err, assert.AnError)
					assert.Nil(t, user)
				},
			)
		},
	)
}
