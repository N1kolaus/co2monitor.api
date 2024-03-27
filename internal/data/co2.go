package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/FMinister/co2monitor-api/internal/validator"
)

type Co2Model struct {
	DB *sql.DB
}

type Co2Data struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	LocationID int64     `json:"location_id"`
	Co2        int       `json:"co2"`
	Temp       float64   `json:"temp"`
	Humidity   int       `json:"humidity"`
}

func (m Co2Model) Insert(co2Data *Co2Data) error {
	query := `
		INSERT INTO co2_data (location_id, co2, temp, humidity)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []any{
		co2Data.LocationID,
		co2Data.Co2,
		co2Data.Temp,
		co2Data.Humidity,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&co2Data.ID, &co2Data.CreatedAt)
}

func (m Co2Model) GetByTimeFrame(id int64, timeFrame time.Duration) ([]*Co2Data, error) {
	query := `
		SELECT id, created_at, location_id, co2, temp, humidity
		FROM co2_data
		WHERE location_id = $1
		AND created_at >= NOW() - $2
		ORDER BY created_at DESC`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		id,
		timeFrame,
	}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var co2Data []*Co2Data

	for rows.Next() {
		var cd Co2Data

		err := rows.Scan(&cd.ID, &cd.CreatedAt, &cd.LocationID, &cd.Co2, &cd.Temp, &cd.Humidity)
		if err != nil {
			return nil, err
		}

		co2Data = append(co2Data, &cd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return co2Data, nil
}

func (m Co2Model) GetLatest(id int64) (*Co2Data, error) {
	query := `
		SELECT id, created_at, location_id, co2, temp, humidity
		FROM co2_data
		WHERE location_id = $1
		ORDER BY created_at DESC
		LIMIT 1`

	var co2Data Co2Data

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&co2Data.ID, &co2Data.CreatedAt, &co2Data.LocationID, &co2Data.Co2, &co2Data.Temp, &co2Data.Humidity)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &co2Data, nil
}

func ValidateCo2Data(v *validator.Validator, co2Data *Co2Data) {
	v.Check(co2Data.LocationID > 0, "location_id", "must greater than zero")

	v.Check(co2Data.Co2 >= 0, "co2", "must be greater than or equal to zero ppm")
	v.Check(co2Data.Co2 < 5000, "co2", "must be less than 5000 ppm")

	v.Check(co2Data.Temp >= -100, "temp", "must be greater than or equal to -100 degrees Celsius")
	v.Check(co2Data.Temp <= 100, "temp", "must be less than or equal to 100 degrees Celsius")

	v.Check(co2Data.Humidity >= 0, "humidity", "must be greater than or equal to zero percent")
	v.Check(co2Data.Humidity <= 100, "humidity", "must be less than or equal to 100 percent")
}
