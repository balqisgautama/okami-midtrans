package router

import (
	"net/http"
)

func MiddlewareCORSOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(*&w).Header().Set("Access-Control-Allow-Origin", "*")
		(*&w).Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization")
		(*&w).Header().Set("Access-Control-Allow-Credentials", "true")
		(*&w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		(*&w).Header().Set("Access-Control-Max-Age", "1209600")
		next.ServeHTTP(w, r)
	})
}
