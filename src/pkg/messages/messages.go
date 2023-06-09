package messages

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)

func InialBanner(){
	banner := figure.NewFigure("MOCKIT", "", true)
	banner.Print()
	fmt.Printf("\n")
}

func PrintParametersErrorMessageAndExit(){
	fmt.Printf("ERROR: \n")
	fmt.Println("Missing parameters")
	fmt.Println("Usage:")
	fmt.Printf("\t mockit VERB ENDPOINT -d RESPONSE_DATA -s STATUS CODE \n")
	fmt.Printf("\n")
	fmt.Printf("\t RESPONSE_DATA and STATUS CODE are optionals \n")
	os.Exit(1)
}

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

func PrintInitialServerLogs(port string, verb string, endpoint string){
	log.Printf(fmt.Sprintf("Starting MOCKIT server at %s port...\n", port))
	log.Println(color.InBold(fmt.Sprintf("\nAvailable endpoint:")))
	fmt.Println(color.InGreen(fmt.Sprintf("\t\t %s http://localhost:%s%s \n", strings.ToUpper(verb),port, endpoint)))
}

func PrintServerRequest(verb string, endpoint string, statusCode int){
	infoLogger := log.New(os.Stdout, color.InBold("INFO: "),log.Ltime)
	upperCaseVerb := strings.ToUpper(verb)
	infoLogger.Println(color.InBlue(fmt.Sprintf("%s %s - %d \n",upperCaseVerb, endpoint, statusCode)))
}

func ShutdownServer(){
	log.Println(color.InBold(fmt.Sprintf("\nClose connections...")))
}

