package example

import "time"

type Example struct {
	Id       int64     `db:"id"`
	FormedAt time.Time `db:"formed_at"`
}
