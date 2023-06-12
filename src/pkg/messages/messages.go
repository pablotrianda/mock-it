package messages

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)
/*
   InialBanner show the header and initial banner in the app 
*/
func InialBanner(){
	banner := figure.NewFigure("MOCKIT", "", true)
	banner.Print()
	fmt.Printf("\n")
}

/*
   PrintParametersErrorMessageAndExit show the error and the correct usage for the app
*/
func PrintParametersErrorMessageAndExit(){
	fmt.Printf("ERROR: \n")
	fmt.Println("Missing parameters")
	fmt.Println("Usage:")
	fmt.Printf("\t mockit VERB ENDPOINT -d RESPONSE_DATA -s STATUS CODE \n")
	fmt.Printf("\n")
	fmt.Printf("\t RESPONSE_DATA and STATUS CODE are optionals \n")
	os.Exit(1)
}

/*
   InvalidVerbErrorMessageAndExit show the incorrect verb and show the list with accepted verbs
   Params:
       - invalidVerb: Invalid verb 
*/
func InvalidVerbErrorMessageAndExit(invalidVerb string){
	fmt.Printf("ERROR: \n")
	fmt.Printf("%v is a invalid http verb \n", strings.ToUpper(invalidVerb))
	fmt.Println("Http verbs supported:")
	fmt.Printf("\t - GET \n")
	fmt.Printf("\t - POST \n")
	fmt.Printf("\t - DELETE \n")
	fmt.Printf("\t - PUT \n")
	fmt.Printf("\t - PATCH \n")
	os.Exit(1)
}

/*
   PrintInitialServerLogs log the server initial status and the addrress for the server mock
   Params:
       - port: Setted server port
       - verb: Http verb
       - endpoint: Endpoint to make the request
*/
func PrintInitialServerLogs(port string, verb string, endpoint string){
	log.Printf(fmt.Sprintf("Starting MOCKIT server at %s port...\n", port))
	log.Println(color.InBold(fmt.Sprintf("\nAvailable endpoint:")))
	fmt.Println(color.InGreen(fmt.Sprintf("\t\t %s http://localhost:%s%s \n", strings.ToUpper(verb),port, endpoint)))
}

/*
   PrintServerRequest Log the server request every time the request is called
   Params:
       - infoLogger: logger to show the info
       - verb: Http verb
       - endpoint: Endpoint to make the request
       - statusCode: Setted status code
*/
func PrintServerRequest(infoLogger *log.Logger,verb string, endpoint string, statusCode int){
	infoLogger.SetPrefix(color.InBold("INFO: "))
	upperCaseVerb := strings.ToUpper(verb)
	infoLogger.Println(color.InBlue(fmt.Sprintf("%s %s - %d \n",upperCaseVerb, endpoint, statusCode)))
}

/*
   ShutdownServer show the message when the server is close
*/
func ShutdownServer(){
	log.Println(color.InBold(fmt.Sprintf("\nClose connections...")))
}

