package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andrefsilveira1/urban/internal/config"
	"github.com/andrefsilveira1/urban/internal/database/scylla"
	"github.com/andrefsilveira1/urban/internal/domain"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
	"github.com/andrefsilveira1/urban/internal/transport/rest"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Starting server...")

	var configpath string
	flag.StringVar(&configpath, "config", "", "Path to config file")
	flag.Parse()

	cfg := loadConfig(configpath)
	db := loadDatabase(cfg.Database)

	userRepository := repository.NewUserRepository(db)
	imageRepository := repository.NewImageRepository(db)
	testRepository := repository.NewTestRepository(db)

	testService := domain.NewTestService(testRepository)
	userService := domain.NewUserService(userRepository)
	imageService := domain.NewImageService(imageRepository)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	var restServer *rest.Server
	g.Go(func() (err error) {
		router := mux.NewRouter().StrictSlash(true)

		rest.NewImageHandler(imageService).Register(router)
		rest.NewUserHandler(userService).Register(router)
		rest.NewTestHandler(testService).Register(router)
		restServer, err = restServer.NewServer(cfg.Server.HTTP, router)
		if err != nil {
			log.Printf("error until new server method: %v", err)
			os.Exit(-1)
		}
		return restServer.Start()
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

	if restServer != nil {
		restServer.Stop(shutdownCtx)
	}

	err := g.Wait()
	if err != nil {
		log.Printf("server shutdown returned an error")
		defer os.Exit(2)
	}
	log.Println("service shutdown")
}

func loadConfig(configPath string) *config.Config {
	if configPath == "" {
		configPath = os.Getenv("APP_CONFIG_PATH")
		if configPath == "" {
			configPath = "./config.yaml"
		}
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Printf("configuration error :%v", err)
		os.Exit(-1)
	}

	return cfg
}

func loadDatabase(cfg *config.Database) *gocql.Session {
	db, err := scylla.Connect(cfg)
	if err != nil {
		log.Printf("load database error: %v", err)
		os.Exit(-1)
	}

	return db
}
