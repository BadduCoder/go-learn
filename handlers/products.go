package handlers

import (
	"golearn/data"
	"log"
	"net/http"
)

type Product struct {
	logger *log.Logger
}

func NewProduct(logger *log.Logger) *Product{
	return &Product{logger}
}

func (product *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		product.GetProducts(rw, r)
		return
	}

	//Catch all other requests
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Product) GetProducts(rw http.ResponseWriter, r *http.Request){
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
