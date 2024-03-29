// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"time"
)

type Account struct {
	ID        int32     `db:"id" json:"id"`
	Owner     string    `db:"owner" json:"owner"`
	Balance   int32     `db:"balance" json:"balance"`
	Currency  *string   `db:"currency" json:"currency"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
