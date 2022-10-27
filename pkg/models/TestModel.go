package models

import (
	"time"
)

type Test struct {
	Id        int
	Title     string
	author_id int
	create_at time.Time
}
