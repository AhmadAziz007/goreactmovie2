package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"log"
	"server/gomovie/helper"
	"server/gomovie/helper/model"
	"server/gomovie/models/web/response"
	"server/gomovie/repository"
)

type MovieService struct {
	MovieRepository *repository.MovieRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewMovieService(movieRepository *repository.MovieRepository, db *sql.DB, validate *validator.Validate) *MovieService {
	return &MovieService{
		MovieRepository: movieRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *MovieService) FindAll(ctx context.Context) []response.MovieResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return nil
	}
	defer helper.CommitOrRollback(tx)

	movieList, err := service.MovieRepository.FindAll(ctx, tx)
	if err != nil {
		log.Printf("Error fetching movies: %v", err)
		return nil
	}

	log.Printf("Fetched movies: %+v\n", movieList)

	return model.ToMovieResponses(movieList)
}
