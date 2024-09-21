package model

import (
	"server/gomovie/models/domain"
	"server/gomovie/models/web/response"
)

func ToMovieResponse(movie domain.Movie) response.MovieResponse {
	return response.MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		Rating:      movie.Rating,
		MPAARating:  movie.MPAARating,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
		MovieGenre:  movie.MovieGenre,
	}
}

func ToMovieResponses(movies []domain.Movie) []response.MovieResponse {
	var movieResponses []response.MovieResponse
	for _, movieList := range movies {
		movieResponses = append(movieResponses, ToMovieResponse(movieList))
	}
	return movieResponses
}
