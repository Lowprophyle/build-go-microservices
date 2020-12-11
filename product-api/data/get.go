package handlers

import (
	"net/http"

	"github.com/lowprophyle/build-go-microservices/product-api/data"
)

// swaggerannotaton [method] [path pattern] [?tag1 tagn] [operation id]
// swagger:route GET /products products listProducts
// Returns a list of product

// GetProducts returns the products from the dummy data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products list")

	// fetch the products from the data store
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
