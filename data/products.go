package data

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	SKU         string `json:"sku"`
	CreationOn  string `json:"-"`
	UpdationOn  string `json:"-"`
	DeletionOn  string `json:"-"`
}

func GetProducts() []*Product {
	pl := productList
	return pl
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
