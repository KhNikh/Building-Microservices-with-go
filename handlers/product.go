package handlers

import (
	"Building-micreoservices-with-go/data"
	"encoding/json"
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
	d := data.GetProducts()

	jd, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	rw.Write(jd)
}
