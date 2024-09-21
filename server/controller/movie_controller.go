package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/gomovie/helper"
	"server/gomovie/models/web"
	"server/gomovie/service"
)

type MovieController struct {
	MovieService *service.MovieService // Perbaiki tipe menjadi pointer
}

func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{
		MovieService: movieService,
	}
}

func (controller *MovieController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieResponse := controller.MovieService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
