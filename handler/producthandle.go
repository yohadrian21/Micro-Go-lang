package handler

import (
	"encoding/json"
	"log"
	"micro/model"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewPorducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// lp := model.GetProduct()
	// d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
	// }
	// rw.Write(d)
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
	}

	if r.Method == http.MethodDelete {
		p.DelProduct(rw, r)
	}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := model.GetProduct()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
	}
	rw.Write(d)

}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	prod := &model.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
	}
	model.AddProduct(prod)
}

func (p *Products) DelProduct(rw http.ResponseWriter, r *http.Request) {

	prod := &model.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
	}
	model.DelProduct(prod)
}
