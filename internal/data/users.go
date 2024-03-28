package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"unicode"

	"github.com/FMinister/co2monitor-api/internal/validator"
)

var (
	ErrDuplicateName = errors.New("duplicate name for user")
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Token     string    `json:"-"`
	Active    bool      `json:"active"`
}

func (m UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users (name, token, active)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`

	args := []interface{}{user.Name, user.Token, user.Active}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_name_key"`:
			return ErrDuplicateName
		default:
			return err

		}
	}

	return nil
}

func (m UserModel) GetByName(name string) (*User, error) {
	query := `
		SELECT id, created_at, updated_at, name, token, active
		FROM users
		WEHRE name = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, name).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
		&user.Token,
		&user.Active)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Update(user *User) error {
	query := `
		UPDATE users
		SET name = $2, token = $3, active = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING updated_at`

	args := []interface{}{
		user.ID,
		user.Name,
		user.Token,
		user.Active,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_name_key"`:
			return ErrDuplicateName
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) Delete(id int64) error {
	query := `
		DELETE FROM users
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

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) < 3, "name", "must be at least 3 bytes long")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(user.Token != "", "token", "must be provided")
	v.Check(len(user.Token) >= 32, "token", "must be at least 32 bytes long")
	v.Check(len(user.Token) <= 72, "token", "must not be more than 72 bytes long")
	v.Check(verifyToken(user.Token), "token", "must contain at least one number, one uppercase letter, one lowercase letter, and one special character")
}

func verifyToken(s string) bool {
	var hasNumber, hasUpperCase, hasLowercase, hasSpecial bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasNumber && hasUpperCase && hasLowercase && hasSpecial
}
