package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
)

// func helloHandler()
func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        d, err := ioutil.ReadAll(r.Body) 
        if err != nil {
            http.Error(rw, "something went wrong", http.StatusBadRequest)
            return
        }
        fmt.Fprintf(rw, "hello %s\n", d)
    })
    fmt.Println("Server running on port 9090")
    http.ListenAndServe(":9090", nil)
}
