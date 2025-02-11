package main

import (
	"log"
	"net/http"

	"github.com/saikumar-neelam/glofox_studio/api/routers"
)

func main() {
	// Setup the router
	router := routers.SetupRouter()

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
