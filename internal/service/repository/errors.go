package repository

import "errors"

var (
	ErrBadRequest     = errors.New("bad request")
	ErrObjectNotFound = errors.New("object not found")
	ErrConflict       = errors.New("conflict: duplicate key value violates unique constraint \"students_pkey\"")
)
