package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/andrefsilveira1/urban/internal/database/scylla"
	"github.com/andrefsilveira1/urban/internal/domain"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
	"github.com/andrefsilveira1/urban/internal/transport/rest"
)

func main() {
	fmt.Println("Starting server...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	session, err := scylla.Connect()
	if err != nil {
		log.Fatalf("Error connecting to Scylla: %v", err)
	}
	defer session.Close()

	// Initialize repositories and services
	userRepository := repository.NewUserRepository(session)
	imageRepository := repository.NewImageRepository(session)

	userService := domain.NewUserService(userRepository)
	imageService := domain.NewImageService(imageRepository)
	rest.Start(3000, userService, imageService)

}
