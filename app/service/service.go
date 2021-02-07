package service

import (
	"github.com/royalbhati/urlshortener/business/shortener"
)

type UrlShortnerService struct {
	Shortener shortener.UrlShortener
}
