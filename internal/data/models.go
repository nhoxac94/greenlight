package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModal
}

func NewModal(db *sql.DB) Models {
	return Models{
		Movies: MovieModal{DB: db},
	}
}
