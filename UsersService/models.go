package main

import (
	"github.com/jmoiron/sqlx"
	"golang.org/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *User) get(db *sqlx.DB) error {
	return db.Get(u, "SELECT name, email FROM users WHERE id=$1", u.ID)
}

func (u *User) update(db *sqlx.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4, u.Name, u.Email, string(hashedPassword), u.ID")
	return err
}

