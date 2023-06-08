package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pablotrianda/mockit/pkg/messages"
)

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	messages.ShutdownServer()
	os.Exit(0)
}

func CreateAndStart(verb string, endpoint string, statusCode int, data map[string]interface{}, port string){
	go gracefulShutdown()

	http.HandleFunc(endpoint,func(w http.ResponseWriter, r *http.Request) {
		messages.PrintServerRequest(verb, endpoint, statusCode)
		// Header and status code
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		// Return a json
		json.NewEncoder(w).Encode(data)
	})
	
	messages.PrintInitialServerLogs(port, verb, endpoint)	
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
