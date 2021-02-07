package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/royalbhati/urlshortener/app/config"
	"github.com/royalbhati/urlshortener/business/shortener"
	"github.com/royalbhati/urlshortener/platform/database"
)

func main() {

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
	if err := run(cfg, os.Args[1]); err != nil {
		log.Println("admin: error:", err)
		os.Exit(1)
	}
}
func run(cfg *config.Config, arg string) error {
	fmt.Println("aaaaaa", arg)
	// Initialize dependencies.
	db, err := database.Open(cfg)

	if err != nil {
		return errors.Wrap(err, "connecting db")
	}
	defer db.Close()

	switch arg {
	case "migrate":
		if err := shortener.Migrate(db); err != nil {
			return errors.Wrap(err, "error applying migrations")
		}
		log.Println("Migrations complete")
		return nil

	case "seed":
		if err := shortener.Seed(db); err != nil {
			return errors.Wrap(err, "error seeding database")
		}
		log.Println("Seed data complete")
		return nil
	}
	return nil
}
