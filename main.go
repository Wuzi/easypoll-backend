package main

import (
	"flag"
	"log"

	"github.com/wunari/easypoll/docs/models"
	"github.com/wunari/easypoll/docs/restapi"
	"github.com/wunari/easypoll/docs/restapi/operations"
	"github.com/wunari/easypoll/docs/restapi/operations/poll"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewEasypollAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	defer func() {
		_ = server.Shutdown()
	}()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// GetGreetingHandler greets the given name,
	// in case the name is not given, it will default to World
	api.PollGetPollsHandler = poll.GetPollsHandlerFunc(
		func(params poll.GetPollsParams) middleware.Responder {
			polls := []*models.Poll{}
			return poll.NewGetPollsOK().WithPayload(polls)
		})

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
