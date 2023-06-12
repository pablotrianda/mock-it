package cli

import (
	"flag"
)

/*
   Cli is an struct to handle and put on a only place the params data
   field:
	- Verb: Http verb
	- Endpoint: Custom end point
	- ResponseData: Reference to json file to build the response 
	- StatusCode: Status code to set on the response
	- ServerPort: Port where the server will be running
*/
type Cli struct{
	Verb string
	Endpoint string
	ResponseData string
	StatusCode string
	ServerPort string
}

/*
   HandleArgs put on a one place all the info passed through the params
   Params:
	- args: []String with params (os.Args)
   Return:
	- Cli: Cli struct with the params loadedd and they default values
*/
func HandleArgs(args []string) Cli {
	cli := Cli{}

	flag.StringVar(&cli.Verb, "v", "get", "Http Verb")
	flag.StringVar(&cli.Endpoint, "e", "/", "End point name")
	flag.StringVar(&cli.ResponseData, "d", "", "Response data")
	flag.StringVar(&cli.StatusCode,"s", "200", "Status code")
	flag.StringVar(&cli.ServerPort,"p", "3000", "Server port")
	flag.Parse()

	return cli
}
