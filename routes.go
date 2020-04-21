package main

import (
	"net/http"
	"github.com/dimfeld/httptreemux"
)

func (app *App) Routes() http.Handler {
	router := httptreemux.New()

	router.GET("/", app.indexHandler)
	router.GET("/*", app.staticHandler)
	router.GET("/healthcheck", app.healthCheckHandler)
	router.NotFoundHandler = app.angularHandler
	return router
}