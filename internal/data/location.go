package data

import (
	"database/sql"
	"time"

	"github.com/FMinister/co2monitor-api/internal/validator"
)

type LocationModel struct {
	DB *sql.DB
}

type Location struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func ValidateLocationData(v *validator.Validator, location *Location) {
	v.Check(location.Name != "", "name", "must be provided")
	v.Check(len(location.Name) > 3, "name", "must be longer than 3 bytes")
	v.Check(len(location.Name) <= 500, "name", "must not be more than 500 bytes long")
}
