package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/royalbhati/urlshortener/cmd/config"
	"github.com/royalbhati/urlshortener/cmd/routes"
	"github.com/royalbhati/urlshortener/platform/database"
	"github.com/royalbhati/urlshortener/platform/redis"
)

func main() {
	log := log.New(os.Stdout, "URL SHORTENER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// Generate our config based on the config supplied
	// by the user in the flags
	cfgFlags, err := config.ParseFlags()
	// log.Println("aaaaa", cfgPath)
	if err != nil {
		log.Println("Config Parse error:", err)
		os.Exit(1)
	}
	cfg, err := config.NewConfig(cfgFlags.Path)
	if err != nil {
		log.Println("Config Parse error:", err)
		os.Exit(1)
	}

	// Run the server
	if err := run(cfg, log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(cfg *config.Config, log *log.Logger) error {

	db, err := database.Open(cfg)
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer db.Close()

	rdb, err := redis.Open(cfg)
	if err != nil {
		return errors.Wrap(err, "connecting to redis db")

	}
	defer rdb.Close()

	api := http.Server{
		Addr:         cfg.Web.Host + ":" + cfg.Web.Port,
		Handler:      routes.API(log, db, rdb),
		ReadTimeout:  cfg.Web.Timeout.Read * time.Second,
		WriteTimeout: cfg.Web.Timeout.Write * time.Second,
	}

	if err := api.ListenAndServe(); err != nil {
		log.Println("ERROR", err)
	}

	return nil
}
