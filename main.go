package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/chenliu1993/webdemo/examples"
	"github.com/goji/httpauth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {
	addr := ":3000"
	authHandler := httpauth.SimpleBasicAuth("username", "password")
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	finalHandler := http.HandlerFunc(examples.FinalFunc)
	loggingHandler := handlers.LoggingHandler(logFile, finalHandler)

	router := mux.NewRouter()
	router.Handle("/", alice.New(authHandler, examples.MiddleWareTwo, examples.MiddleWareOne).Then(loggingHandler))
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	if err := server.ServeTLS(listener, "key.pem", "cert.pem"); err != nil {
		log.Fatal(err)
	}
	return
}
