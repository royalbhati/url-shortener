package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Router struct {
	mux *chi.Mux
	log *log.Logger
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func NewRouter(log *log.Logger) *Router {
	return &Router{
		mux: chi.NewRouter(),
		log: log,
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func (router *Router) Handle(method string, path string, customHandler Handler) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := customHandler(w, r); err != nil {
			Response(w, err.Error(), http.StatusInternalServerError)
			router.log.Printf("Unhandled error: %+v", err)
		}
	}

	router.mux.MethodFunc(method, path, handler)
}
