package routes

import (
	"log"
	"net/http"

	"github.com/royalbhati/urlshortener/app/service"
	"github.com/royalbhati/urlshortener/business/shortener"
	"github.com/royalbhati/urlshortener/platform/router"
)

func API(logger *log.Logger) http.Handler {

	us := service.UrlShortnerService{
		Shortener: shortener.New(logger),
	}

	router := router.NewRouter(logger)
	router.Handle(http.MethodGet, "/v1/test", us.Test)
	return router
}
