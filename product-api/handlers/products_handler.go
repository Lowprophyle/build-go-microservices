// Package handlers of Product API
//
// Documentation for Product API
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lowprophyle/build-go-microservices/product-api/product-api/data"

	"github.com/gorilla/mux"
)

type productResponse struct {
	Body []data.Product
}

type Products struct {
	logger *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
// 	p.logger.Println("Handle GET Products list")
// 	lp := data.GetProducts()
// 	err := lp.ToJSON(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
// 	}
// }

func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle Product update via PUT")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.logger.Println("PUT UpdateProduct for this ID:", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle new Product via POST")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.logger.Printf("Prod: %#v", &prod)
	data.AddProduct(&prod)
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		p.logger.Println("req body", r.Body)
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.logger.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
