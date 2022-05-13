package restful

import (
	"net/http"
	"testing"

	"github.com/emicklei/go-restful/v3"
)

func TestRestfulContainer(t *testing.T) {
	container := restful.NewContainer()

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Filter(func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		t.Logf("[print-filter] %s, %s\n", req.Request.Method, req.Request.URL)
		chain.ProcessFilter(req, resp)
	})

	ws.Path("/users")
	ws.Route(ws.GET("/{userId}").To(func(req *restful.Request, resp *restful.Response) {
		id := req.PathParameter("userId")
		resp.WriteAsJson(map[string]string{
			"id": id,
		})
	}))
	container.Add(ws)

	server := &http.Server{Addr: ":8080", Handler: container}
	server.ListenAndServe()
}
