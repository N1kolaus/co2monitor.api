package data

import (
	"database/sql"
	"errors"
)

var (
	ErrEditConflict   = errors.New("edit conflict")
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	CO2        Co2Model
	Location   LocationModel
	Permission PermissionModel
	Tokens     TokensModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		CO2:        Co2Model{DB: db},
		Location:   LocationModel{DB: db},
		Permission: PermissionModel{DB: db},
		Tokens:     TokensModel{DB: db},
	}
}
