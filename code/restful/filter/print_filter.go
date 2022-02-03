package filter

import (
	"fmt"
	"net/http"
)

func WithPrintFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("print...")
		handler.ServeHTTP(w, req)
	})
}
