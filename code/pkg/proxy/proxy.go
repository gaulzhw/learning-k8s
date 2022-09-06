package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// modify request
		req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("X-Resoponse-Proxy", "Simple-Reverse-Proxy")
		return nil
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v", err)
	}

	return proxy, nil
}
