package middleware

import (
	"net/http"
	"strings"
)

const (
	prefixUserAgent = "User-Agent"
)

func Headers(next http.Handler) http.Handler {
	if lgr == nil {
		return next
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		headLog := make(map[string]interface{})

		for name, values := range r.Header {
			if strings.HasPrefix(name, prefixUserAgent) {
				headLog[name] = values
			}

		}

		if len(headLog) > 0 {
			tid := GetRequestID(r.Context())
			lgr.Println("Headers", tid, headLog)
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
