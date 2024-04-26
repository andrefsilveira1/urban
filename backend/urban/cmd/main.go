package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andrefsilveira1/urban/internal/config"
	"github.com/andrefsilveira1/urban/internal/transport/rest"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Starting server...")

	cfg := &config.ServerHTTP{Host: "localhost", Port: 8080}

	// Gracefully shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	var rest *rest.Server
	g.Go(func() (err error) {
		router := mux.NewRouter().StrictSlash(true)
		rest, err = rest.NewServer(cfg, router)

		return rest.Start()
	})

	log.Println("service started")

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Println("shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if rest != nil {
		rest.Stop(shutdownCtx)
	}

	err := g.Wait()
	if err != nil {
		log.Printf("server shutdown returned an error")
		defer os.Exit(2)
	}
	log.Println("service shutdown")
}
