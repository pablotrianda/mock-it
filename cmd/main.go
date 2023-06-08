package main

import (
	"os"
	"strconv"

	c "github.com/pablotrianda/mockit/pkg/cli"
	"github.com/pablotrianda/mockit/pkg/messages"
	"github.com/pablotrianda/mockit/pkg/response"
	"github.com/pablotrianda/mockit/pkg/server"
	"github.com/pablotrianda/mockit/pkg/validations"
)

func main(){
	cli := c.HandleArgs(os.Args)
	err := validations.ValidateVerb(cli.Verb)
	if err != nil {
		messages.InvalidVerbErrorMessageAndExit(cli.Verb)
	}
	
	// Validate and parse data
	responseData, err := response.Process(cli.ResponseData)
	statusCode, err := strconv.Atoi(cli.StatusCode)
	endpoint := validations.FormatEndpoint(cli.Endpoint)

	// Message banner
	messages.InialBanner()

	// Start server
	server.CreateAndStart(cli.Verb, endpoint, statusCode, responseData, cli.ServerPort)
}
