package middleware

import (
	"github.com/suaas21/graphql-dummy/api/response"
	"net/http"
)

func AppKeyChecker(appKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("Application-Key")
			if len(key) == 0 {
				response.ServeJSON(w, nil, http.StatusUnauthorized)
				return
			}
			if key != appKey {
				response.ServeJSON(w, nil, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
