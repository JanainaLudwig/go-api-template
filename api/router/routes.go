package router

import (
	"github.com/julienschmidt/httprouter"
	"go-api-template/api/handlers"
	"net/http"
)

func Routes() http.Handler {
	router := httprouter.New()
	router.GET("/", handlers.Index)

	return router
}