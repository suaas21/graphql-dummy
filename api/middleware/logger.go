package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/suaas21/graphql-dummy/logger"
	"net/http"
)

// Logger returns a request logging middleware
func Logger(lgr logger.Logger) Middleware {
	if lgr == nil {
		return func(h http.Handler) http.Handler {
			return h
		}
	}
	return middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: lgr})
}

var RequestID = middleware.RequestID

var GetRequestID = middleware.GetReqID
