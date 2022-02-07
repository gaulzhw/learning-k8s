package restful

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
)

func print(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	fmt.Printf("[print-filter] %s, %s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

func findUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("userId")
	resp.WriteAsJson(map[string]string{
		"id": id,
	})
}

func findMessage(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("messageId")
	resp.WriteAsJson(map[string]string{
		"id": id,
	})
}
