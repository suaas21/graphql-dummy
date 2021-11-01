package middleware

// ContextKey hold the key of a context
type ContextKey string

// List of contexts
const (
	UserContext ContextKey = "user"
)

/*
func GetUser(r *http.Request) *service.UserData {
	v := r.Context().Value(UserContext)
	if v == nil {
		panic(errors.New("middleware: GetUser called without calling auth middleware prior"))
	}
	u, _ := v.(*service.UserData)
	return u
}

// Auth returns authentication middleware
func Auth(auth *service.Auth) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tok := r.Header.Get("Authorization")
			if !strings.HasPrefix(tok, "Bearer ") {
				resp.ServeUnauthorized(w, r, errors.New("unauthorized"))
				return
			}
			tok = strings.TrimSpace(strings.TrimPrefix(tok, "Bearer "))
			if tok == "" {
				resp.ServeUnauthorized(w, r, errors.New("unauthorized"))
				return
			}
			_, u, err := auth.Check(tok)
			if err != nil {
				if err == service.ErrUserNotFound ||
					err == service.ErrUserDisabled {
					resp.ServeUnauthorized(w, r, errors.New("unauthorized"))
					return
				}
				resp.ServeInternalServerError(w, r, err)
				return
			}
			if u == nil {
				resp.ServeUnauthorized(w, r, errors.New("unauthorized"))
				return
			}

			ctx := context.WithValue(r.Context(), UserContext, u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}*/
