package main

import (
	"Building-micreoservices-with-go/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	//create a new server mux and register the handlers
	sm := mux.NewRouter()
	// sm.Handle("/", hh)

	// gh := handlers.NewBye(l)
	// sm.Handle("/goodbye", gh)

	ph := handlers.NewProduct(l)
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.JsonValidationMiddleware)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.JsonValidationMiddleware)
	// sm.Handle("/products/", ph)
	// Creating a custom server and
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		fmt.Println("Server running on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	sig := <-sigchan
	l.Println("Recieved terminate, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
