package handlers

import (
	"Building-micreoservices-with-go/data"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling Post Request")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(id, prod)
	if err == data.ProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Product) JsonValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.l.Println("cp---1")
		prod := data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, &prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
