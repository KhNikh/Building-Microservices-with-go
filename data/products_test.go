package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		ID:    4,
		Name:  "tea",
		Price: 2,
		SKU:   "abc-def-ghi",
	}
	err := p.ValidateProduct()
	if err != nil {
		t.Fatal(err)
	}
}
