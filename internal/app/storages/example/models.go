package example

import "time"

type Example struct {
	ID       int64     `db:"id"`
	FormedAt time.Time `db:"formed_at"`
}
