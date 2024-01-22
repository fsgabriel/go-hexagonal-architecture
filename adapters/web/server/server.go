package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/fsgabriel/go-hexagonal-architecture/adapters/web/server/handler"
	"github.com/fsgabriel/go-hexagonal-architecture/application"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (s *WebServer) Server() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, s.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
