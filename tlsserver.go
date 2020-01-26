package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/chenliu1993/webdemo/examples"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {
	caCrtPath := "ca.crt"
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(caCrtPath)
	if err != nil {
		log.Fatal(err)
	}
	pool.AppendCertsFromPEM(caCrt)

	addr := ":3001"
	// authHandler := httpauth.SimpleBasicAuth("cliu2", "cliu2")
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	finalHandler := http.HandlerFunc(examples.FinalFunc)
	loggingHandler := handlers.LoggingHandler(logFile, finalHandler)

	router := mux.NewRouter()
	// router.Handle("/", alice.New(authHandler, examples.MiddleWareTwo, examples.MiddleWareOne).Then(loggingHandler))
	router.Handle("/", alice.New(examples.MiddleWareTwo, examples.MiddleWareOne).Then(loggingHandler))
	server := &http.Server{
		Addr:    addr,
		Handler: router,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	if err := server.ServeTLS(listener, "server.crt", "server.key"); err != nil {
		log.Fatal(err)
	}
	return
}
