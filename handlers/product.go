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
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {
		p.updateProduct(rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling Post Request")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	data.AddProduct(prod)
}

func (p *Product) updateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Put Request")
	url := r.URL.Path
	fmt.Println("url: ", url)
}
