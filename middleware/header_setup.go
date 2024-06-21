package middleware

import "net/http"

func HeaderSetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header, Since we are always returning JSON
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
