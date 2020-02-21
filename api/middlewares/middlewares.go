package middlewares

import (
	"errors"
	"net/http"

	"rest/api/auth"
	"rest/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "http://local.detik.com:8080")
		header.Add("Content-Type", "application/json")
		header.Add("Access-Control-Allow-Credentials", "true")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		next.ServeHTTP(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://local.detik.com:8080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	}
}
