package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"gt=0"`
	SKU         string `json:"sku" validate:"required,sku"`
	CreationOn  string `json:"-"`
	UpdationOn  string `json:"-"`
	DeletionOn  string `json:"-"`
}

func (p *Product) ValidateProduct() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	fmt.Println("final checkpoint")
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	fmt.Println("sku validation")
	sku := fl.Field().String()
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	result := re.FindAllString(sku, -1)
	if len(result) != 1 {
		return false
	}
	return true

}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)

}

func (p *Product) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products {
	lp := productList
	return lp
}

func AddProduct(prod *Product) {
	prod.ID = getNextId()
	productList = append(productList, prod)
}

func UpdateProduct(id int, prod *Product) error {
	_, idx, err := findProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	productList[idx] = prod
	return nil
}

func getNextId() int {
	return len(productList) + 1
}

var ProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "RDPD",
		Description: "A real life story",
		Price:       280,
		SKU:         "nikh111",
		CreationOn:  "2023-09-21 12:34:56 UTC",
		UpdationOn:  "2023-09-22 12:34:56 UTC",
		DeletionOn:  "2023-09-23 12:34:56 UTC",
	},
	&Product{
		ID:          2,
		Name:        "The Chronicles of Elantris",
		Description: "A fictional story",
		Price:       320,
		SKU:         "nikh112",
		CreationOn:  "2023-09-21 12:34:56 UTC",
		UpdationOn:  "2023-09-22 12:34:56 UTC",
		DeletionOn:  "2023-09-23 12:34:56 UTC",
	},
}
