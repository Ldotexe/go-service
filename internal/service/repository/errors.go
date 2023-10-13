package repository

import "errors"

var (
	ErrObjectNotFound = errors.New("object not found")
	ErrConflict       = errors.New("conflict: duplicate key value violates unique constraint \"students_pkey\" ")
)
