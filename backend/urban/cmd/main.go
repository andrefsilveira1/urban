package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andrefsilveira1/urban/internal/database/scylla"
	"github.com/andrefsilveira1/urban/internal/domain"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
	"github.com/andrefsilveira1/urban/internal/transport/rest"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Starting server...")
	// var path string // database variables will got be here

	// put repositories

	// Shutdown

	// Create Dependency injection here {}
	// {}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	session, err := scylla.Connect()
	if err != nil {
		log.Fatalf("Error connecting to Scylla: %v", err)
	}
	defer session.Close()
	userRepository := repository.NewScyllaUserRepository(session)
	userService := domain.NewUserService(userRepository)

	g.Go(func() (err error) {
		fmt.Println("Server started")
		err = rest.Start(3000, userService)
		// Start rest server here
		if err != nil {
			panic(err)
		}
		return
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Println("Shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if shutdownCtx != nil {
		fmt.Println("Shutdown started")
		// stop
	}

	err = g.Wait()
	if err != nil {
		log.Printf("Server shutdown returned an error")
		defer os.Exit(2)
	}

	log.Println("Shutdown")

}

// load config

// load database
