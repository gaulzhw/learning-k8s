package restful

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

/*
func multiContainers() {
	// add to default container
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// container 2
	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/hello").To(hello2))
	container2.Add(ws2)
	server := &http.Server{Addr: ":8081", Handler: container2}
	log.Fatal(server.ListenAndServe())
}
*/

func StartContainer() {
	wsContainer := restful.NewContainer()

	userWS := new(restful.WebService)
	userWS.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	userWS.Filter(print)
	userWS.Path("/users")
	userWS.Route(userWS.GET("/{userId}").To(findUser))
	wsContainer.Add(userWS)

	msgWS := new(restful.WebService)
	msgWS.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	msgWS.Filter(print)
	msgWS.Path("/messages")
	msgWS.Route(msgWS.GET("/{messageId}").To(findMessage))
	wsContainer.Add(msgWS)

	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	server.ListenAndServe()
}
