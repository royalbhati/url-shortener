package routes

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/royalbhati/urlshortener/app/shortener"
	"github.com/royalbhati/urlshortener/cmd/service"
	"github.com/royalbhati/urlshortener/platform/router"
)

func API(logger *log.Logger, db *sqlx.DB, redis *redis.Client) http.Handler {

	us := service.UrlShortnerService{
		Shortener: shortener.New(logger, db, redis),
	}

	router := router.NewRouter(logger)
	router.Handle(http.MethodPost, "/v1/short", us.GetShort)
	router.Handle(http.MethodGet, "/v1/long/{id}", us.GetLong)

	return router
}
