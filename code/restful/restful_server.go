package restful

import (
	"net/http"

	"github.com/emicklei/go-restful"

	"github.com/gaulzhw/learning-k8s/restful/filter"
)

func StartServer() {
	container := restful.NewContainer()

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Path("/users")
	ws.Route(ws.GET("/{userId}").To(findUser))

	container.Add(ws)

	directorHandler := &director{
		container: container,
	}

	apiServerHandler := &apiServer{
		director: buildHandlerChain(directorHandler),
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: apiServerHandler,
	}
	srv.ListenAndServe()
}

func buildHandlerChain(handler http.Handler) http.Handler {
	handler = filter.WithLogFilter(handler)
	handler = filter.WithPrintFilter(handler)
	return handler
}

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
