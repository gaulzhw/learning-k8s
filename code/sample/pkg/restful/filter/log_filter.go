package filter

import (
	"log"
	"net/http"
)

func WithLogFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("log...")
		handler.ServeHTTP(w, req)
	})
}
