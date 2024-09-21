package create

import "time"

type MovieCreateRequest struct {
	Title       string         `validate:"required,min=1,max=100" json:"title"`
	Description string         `validate:"required,min=1,max=100" json:"description"`
	Year        int            `validate:"required"               json:"year"`
	ReleaseDate time.Time      `validate:"required"               json:"release_date"`
	Runtime     int            `validate:"required"               json:"runtime"`
	Rating      int            `validate:"required"               json:"rating"`
	MPAARating  string         `validate:"required,min=1,max=100" json:"mpaa_rating"`
	CreatedAt   time.Time      `validate:"required"               json:"created_at"`
	UpdatedAt   time.Time      `validate:"required"               json:"updated_at"`
	MovieGenre  map[int]string `validate:"required"               json:"genres"`
}
