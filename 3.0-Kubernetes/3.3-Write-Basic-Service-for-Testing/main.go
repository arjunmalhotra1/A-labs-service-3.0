package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
	We will set up a service to get our basic Kubernetes environment up & running.
	1. Start & stop the service cleanly.
	2. Logging.
*/

// Package level variable.
var build = "develop"

func main() {
	log.Println("Starting Service", build)
	defer log.Println("service ended")

	shutdown := make(chan os.Signal, 1)
	// SIGNINT is for "ctrl+C" and SIGTERM is what kubernetes is going to send for
	// shutdown.
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Stopping service ")
}
