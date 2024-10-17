package main

import (
	"fmt"

	"github.com/ale8k/hexagonal-architecture-in-go-example/internal/adaptors"
	"github.com/ale8k/hexagonal-architecture-in-go-example/internal/core"
)

func main() {
	// Create Driven Adaptors
	fmt.Println("Setting up driven adaptors")
	dbAdaptor := adaptors.NewDbAdaptor()

	// Create Core Services (utilising driven adaptors)
	fmt.Println("Setting up core services")
	mailService := core.NewMailService(dbAdaptor)

	// Create Driving Adaptors (utilising core services)
	fmt.Println("Setting up driving adaptors")
	httpAdaptor := adaptors.NewHttpAdaptor(mailService)

	httpAdaptor.Run()
}
