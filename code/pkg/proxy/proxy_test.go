package proxy

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPProxy(t *testing.T) {
	proxy, err := NewProxy("http://www.baidu.com")
	assert.NoError(t, err)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		proxy.ServeHTTP(w, req)
	})
	http.ListenAndServe(":8080", nil)
}
