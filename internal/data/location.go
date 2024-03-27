package data

import (
	"context"
	"database/sql"
	"errors"
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

func (m LocationModel) Insert(location *Location) error {
	query := `
		INSERT INTO locations (name)
		VALUES ($1)
		RETURNING id, created_at, updated_at`

	args := []any{
		location.Name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&location.ID, &location.CreatedAt, &location.UpdateAt)
}

func (m LocationModel) GetAll() ([]*Location, error) {
	query := `
		SELECT id, created_at, updated_at, name
		FROM locations
		ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var locations []*Location

	for rows.Next() {
		var location Location

		err := rows.Scan(&location.ID, &location.CreatedAt, &location.UpdateAt, &location.Name)
		if err != nil {
			return nil, err
		}

		locations = append(locations, &location)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}

func (m LocationModel) Get(id int64) (*Location, error) {
	query := `
		SELECT id, created_at, updated_at, name
		FROM locations
		WHERE id = $1`

	var location Location

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&location.ID, &location.CreatedAt, &location.UpdateAt, &location.Name)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &location, nil
}

func (m LocationModel) Update(location *Location) error {
	query := `
		UPDATE locations
		SET name = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING updated_at`

	args := []any{
		location.ID,
		location.Name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&location.UpdateAt)
}

func (m LocationModel) Delete(id int64) error {
	query := `
		DELETE FROM locations
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
