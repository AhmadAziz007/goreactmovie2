package domain

import "time"

type Movie struct {
	ID          int
	Title       string
	Description string
	Year        int
	ReleaseDate time.Time
	Runtime     int
	Rating      int
	MPAARating  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	MovieGenre  map[int]string
}
