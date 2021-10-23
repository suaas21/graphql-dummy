package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/suaas21/graphql-dummy/api/middleware"
	"github.com/suaas21/graphql-dummy/logger"
)

var lgr logger.Logger

func SetLogger(l logger.Logger) {
	lgr = l
}
func NewRouter(baCtrl *BookAuthorController) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger(lgr))
	router.Use(middleware.Headers)
	router.Use(middleware.Cors())
	router.Use(chimiddleware.Timeout(30 * time.Second))

	router.NotFound(NotFoundHandler)
	router.MethodNotAllowed(MethodNotAllowed)

	router.Route("/", func(r chi.Router) {
		r.Get("/ok", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
			return
		})
		r.HandleFunc("/graphql", baCtrl.BookAuthorController)
	})
	return router
}

// NotFoundHandler handles when no routes match
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// MethodNotAllowed handles when no routes match
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
