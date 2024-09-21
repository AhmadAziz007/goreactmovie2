package domain

import (
	"time"
)

type MovieGenre struct {
	ID        int
	MovieID   int
	GenreID   int
	Genre     Genre
	CreatedAt time.Time
	UpdatedAt time.Time
}
