package models

import "time"

type User struct {
	ID         int64  `db:"id"            json:"id"`
	First_name string `db:"first_name"    json:"first_name"`
	Last_name  string `db:"last_name"     json:"last_name"`
	Username   string `db:"username"      json:"username"`
	Email      string `db:"email"         json:"email"`
	// Password  string `db:"password" json:"-"`
	DateJoined time.Time `db:"registered_at" json:"dateJoined"`
	// Karma     int    `db:"karma" json:"karma"`
}
