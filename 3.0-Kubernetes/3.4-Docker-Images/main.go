package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
	Before we can run this code in Kubernetes, we need a docker image that can run the binary.
	We need to create a docker image, including the binary that we want to run & then
	Kubernetes can take that image construct a container from it.

	First step is to add some support to create a docker image.
	We will put all the configuration in a folder called, "Zarf".
	We will talk about the project folders later.
	"Zarf" is a sleeve that we put over a hot cup of coffee.
	"Zarf" is like a sleeve that we put over our container so that we don't burn ourselves.
	"Zarf" layer will contain all the configuration that we need.
	We will add a "docker" folder under "zarf".

*/

var build = "develop"

func main() {
	log.Println("Starting Service", build)
	defer log.Println("service ended")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Stopping service ")
}
