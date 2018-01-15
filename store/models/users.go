package models

import (
	"github.com/nicewrk/design-brain-api/store"
)

// API represents the methods that are necessary to implement.
type API interface {
	Insert(store.API) error
	Select(store.API) (User, error)
}

// User implements API.
type User struct {
	UID        string `json:"uid"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	IsVerified bool   `json:"is_verified"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// Insert INSERTs a single user record.
func (u *User) Insert(c store.API) error {
	SQL := `
		INSERT INTO users
		(email, username)
		VALUES
		($1, $2);
	`

	stmt, tx, err := c.Prepare(SQL)
	if err != nil {
		return err
	}

	return c.Execute(stmt, tx, u.Email, u.Username)
}

// Select SELECTs a single user record.
func (u *User) Select(c store.API) (*User, error) {
	SQL := `
		SELECT *
		FROM users
		WHERE username = $1;
	`

	stmt, tx, err := c.Prepare(SQL)
	if err != nil {
		return u, err
	}

	err = stmt.QueryRow(u.Username).Scan(&u.UID, &u.Email, &u.Username, &u.IsVerified, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, c.Rollback(tx, err)
	}
	return u, tx.Commit()
}
