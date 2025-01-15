package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting Broker Service on Port %s\n", webPort)

	// define HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.router(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
