package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Starting server...")
	// var path string // database variables will got be here

	// put repositories

	// Shutdown

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() (err error) {
		fmt.Println("Server started")
		// Start rest server here
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

	err := g.Wait()
	if err != nil {
		log.Printf("Server shutdown returned an error")
		defer os.Exit(2)
	}

	log.Println("Shutdown")

}
