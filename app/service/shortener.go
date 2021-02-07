package service

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/royalbhati/urlshortener/business/shortener"
	"github.com/royalbhati/urlshortener/platform/router"
)

func (u *UrlShortnerService) GetShort(w http.ResponseWriter, r *http.Request) error {
	var ni shortener.NewItem
	if err := router.Decode(r, &ni); err != nil {
		return errors.Wrapf(err, "unable to decode payload")
	}

	data, err := u.Shortener.CreateShortURL(r.Context(), ni)

	//TODO: Handle Error Responses
	if err != nil {
		return errors.Wrapf(err, "Creating URL %q")
	}

	router.Response(w, data, http.StatusOK)
	return nil
}

func (u *UrlShortnerService) GetLong(w http.ResponseWriter, r *http.Request) error {
	shortUrl := chi.URLParam(r, "id")
	data, err := u.Shortener.GetBigURL(r.Context(), shortUrl)
	if err != nil {
		return errors.Wrapf(err, "Retrieving URL %q")
	}
	router.Response(w, data, http.StatusOK)
	return nil
}
