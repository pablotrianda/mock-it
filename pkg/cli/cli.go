package cli

import (
	"flag"
)

type Cli struct{
	Verb string
	Endpoint string
	ResponseData string
	StatusCode string
	ServerPort string
}

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
