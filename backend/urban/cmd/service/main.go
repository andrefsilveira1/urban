package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
	Config *config.ServerHTTP
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
