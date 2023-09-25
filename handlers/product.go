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

// func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProduct(rw, r)
// 		return
// 	} else if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	} else if r.Method == http.MethodPut {
// 		p.l.Println("PUT", r.URL.Path)
// 		urlPattern := "/([0-9]+)"
// 		re := regexp.MustCompile(urlPattern)
// 		matchGroup := re.FindAllStringSubmatch(r.URL.Path, -1)
// 		if len(matchGroup) != 1 {
// 			p.l.Println("Invalid URI, more than one capture group")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		} else if len(matchGroup[0]) != 2 {
// 			p.l.Println("Invalid URI, more than one Id")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		id_string := matchGroup[0][1]
// 		id, err := strconv.Atoi(id_string)

// 		if err != nil {

// 		}
// 		p.updateProduct(rw, r, id)
// 		return

// 	}
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *Product) GetProduct(rw http.ResponseWriter, r *http.Request) {
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

func (p *Product) updateProduct(rw http.ResponseWriter, r *http.Request, id int) {
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
}
