package filter

import (
	"log"
	"net/http"
)

func WithPrintFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("print...")
		handler.ServeHTTP(w, req)
	})
}
