package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID				int		`json:"id"`
	Name			string	`json:"name"`
	Description 	string	`json:"description"`
	Price 			float32	`json:"price"`
	SKU 			string	`json:"sku"`
	CreatedOn		string	`json:"-"`
	UpdatedOn		string	`json:"-"`
	DeletedOn		string	`json:"-"`
}

type Products []*Product

func (products *Products) ToJSON(w io.Writer) error{
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

func GetProducts() Products{
	return productList
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Cappuccino",
		Description: "With Sugar",
		Price:       2.99,
		SKU:         "Random SKU",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Frappe",
		Description: "Without Sugar",
		Price:       2.29,
		SKU:         "Other SKU",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
