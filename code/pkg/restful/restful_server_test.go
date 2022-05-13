package restful

import (
	"net/http"
	"testing"

	"github.com/emicklei/go-restful/v3"
)

type apiServer struct {
	director http.Handler
}

func (h *apiServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.director.ServeHTTP(w, req)
}

type director struct {
	container *restful.Container
}

func (h *director) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.container.Dispatch(w, req)
}

func TestRestfulServer(t *testing.T) {
	container := restful.NewContainer()

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Path("/users")
	ws.Route(ws.GET("/{userId}").To(func(req *restful.Request, resp *restful.Response) {
		id := req.PathParameter("userId")
		resp.WriteAsJson(map[string]string{
			"id": id,
		})
	}))
	container.Add(ws)

	apiServerHandler := &apiServer{
		director: func() http.Handler {
			var handler http.Handler
			handler = &director{
				container: container,
			}
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				t.Log("log...")
				handler.ServeHTTP(w, req)
			})
			handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				t.Log("print...")
				handler.ServeHTTP(w, req)
			})
			return handler
		}(),
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: apiServerHandler,
	}
	srv.ListenAndServe()
}
