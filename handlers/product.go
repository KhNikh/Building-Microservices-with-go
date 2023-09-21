package handlers

import (
	"Building-micreoservices-with-go/data"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	err := data.ToJson(rw)
	if err != nil {
		fmt.Println(err)
	}
}
