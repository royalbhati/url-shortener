package service

import (
	"net/http"

	"github.com/royalbhati/urlshortener/business/shortener"
)

type UrlShortnerService struct {
	Shortener shortener.UrlShortener
}

func (u *UrlShortnerService) Test(w http.ResponseWriter, r *http.Request) error {
	u.Shortener.TestServ()
	return nil
}
