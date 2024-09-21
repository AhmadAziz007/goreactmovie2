package domain

import "time"

type Genre struct {
	ID        int
	GenreName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
