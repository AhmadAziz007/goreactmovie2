package app

import (
	"github.com/julienschmidt/httprouter"
	"server/gomovie/controller"
	"server/gomovie/exception"
)

func NewRouter(movieController *controller.MovieController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/movies", movieController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
