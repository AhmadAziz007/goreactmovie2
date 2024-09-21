package repository

import (
	"context"
	"database/sql"
	"log"
	"server/gomovie/models/domain"
)

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (repository *MovieRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Movie, error) {
	movieQuery := "SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at FROM movies ORDER BY id"

	rows, err := tx.QueryContext(ctx, movieQuery)
	if err != nil {
		log.Printf("Error executing movie query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var movieList []domain.Movie

	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning movie: %v", err)
			continue
		}

		// Query genres for each movie
		genreQuery := "SELECT mg.id, mg.movie_id, mg.genre_id, g.genre_name FROM movies_genres mg LEFT JOIN genres g ON (g.id = mg.genre_id) WHERE mg.movie_id = $1"
		genreRows, err := tx.QueryContext(ctx, genreQuery, movie.ID)
		if err != nil {
			log.Printf("Error executing genre query: %v", err)
			return nil, err
		}
		defer genreRows.Close()

		genres := make(map[int]string)

		for genreRows.Next() {
			var mg domain.MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName)
			if err != nil {
				log.Printf("Error scanning genre: %v", err)
				continue
			}

			// Add genres to map
			genres[mg.GenreID] = mg.Genre.GenreName
		}
		movie.MovieGenre = genres
		movieList = append(movieList, movie)
		log.Printf("Fetched movie: %+v", movie)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after looping through rows: %v", err)
		return nil, err
	}

	return movieList, nil
}
