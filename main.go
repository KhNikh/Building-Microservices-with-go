package main

import (
	"Building-micreoservices-with-go/handlers"
	"fmt"
	"log"
	"os"

	// "log"
	"net/http"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	gh := handlers.NewBye(l)
	sm.Handle("/goodbye", gh)
	fmt.Println("Server running on port 9090")
	http.ListenAndServe(":9090", sm)
}
