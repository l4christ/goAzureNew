package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Println("Timer trigger function starting.")

	// Run the function immediately for testing purposes
	executeTimerFunction()

	// Set up a web server to keep the process alive (needed for Azure Functions runtime)
	http.HandleFunc("/", timerTriggerHandler)
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func timerTriggerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Timer trigger function executed at:", time.Now())
	executeTimerFunction()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Timer trigger function executed successfully"))
}

func executeTimerFunction() {
	log.Println("Hello, this is your scheduled function running!")
}
