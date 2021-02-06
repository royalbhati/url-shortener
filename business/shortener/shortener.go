package shortener

import (
	"log"
)

type UrlShortener struct {
	log *log.Logger
}

// New constructs a Product for api access.
func New(log *log.Logger) UrlShortener {
	return UrlShortener{
		log: log,
	}
}

func (u UrlShortener) CreateShortURL() error {
	u.log.Println("BUSINESS LAYER")
	return nil
}
