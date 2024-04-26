package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andrefsilveira1/urban/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
	Config *config.ServerHTTP
}

func NewServer(cfg *config.ServerHTTP, router *mux.Router) *Server {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It's alive!")
	}).Methods(http.MethodGet)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
		Config: cfg,
	}
}

func (s *Server) Start() error {
	var err error

	log.Printf("Starting HTTP server at '%s:%d'\n", s.Config.Host, s.Config.Port)

	if s.Config.UseHTTPS {
		log.Println("SSL certificate enabled")
		certPath := s.Config.CertPath
		err = s.Server.ListenAndServeTLS(
			fmt.Sprintf("%s/server.crt", certPath),
			fmt.Sprintf("%s/server.key", certPath),
		)
	} else {
		log.Println("SSL certificate disabled")
		err = s.Server.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		log.Printf("Unable to start HTTP server: %+v\n", err)
		return err
	}

	return nil
}

func main() {
	log.Println("Service starting...")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "It is Alive!!!!")
	})

	addr := ":8080"
	log.Printf("Starting http server at '%s' \n", addr)
	http.ListenAndServe(addr, router)

	log.Println("Service shutdown")
}
