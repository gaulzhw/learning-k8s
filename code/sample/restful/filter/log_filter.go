package filter

import (
	"fmt"
	"net/http"
)

func WithLogFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("log...")
		handler.ServeHTTP(w, req)
	})
}
