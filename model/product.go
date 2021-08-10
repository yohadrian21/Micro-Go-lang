package model

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetProduct() []*Product {
	return Prodlist
}

// func (p *Product) AddProduct(r io.Reader) error {
// 	e := json.NewDecoder(r)
// 	return e.Decode(p)
// 	//Prodlist = append(Prodlist, e)
// }
type Products []*Product

func (p *Product) ToJson(r io.Writer) error {
	e := json.NewEncoder(r)
	return e.Encode(p)
	//Prodlist = append(Prodlist, e)
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
	//Prodlist = append(Prodlist, e)
}
func AddProduct(p *Product) {
	cnt := 0

	for i := 0; i < len(Prodlist); i++ {
		if Prodlist[i].ID == p.ID {
			Prodlist[i].Name = p.Name
			cnt = 1
		}
	}
	if cnt == 0 {
		p.ID = getNextId()
		Prodlist = append(Prodlist, p)
	}
	//Prodlist = append(Prodlist, e)
}

func DelProduct(p *Product) {

	for i := 0; i < len(Prodlist); i++ {
		if Prodlist[i].ID == p.ID {
			remove(Prodlist, i)
		}
	}

	// Prodlist,_=Prodlist.Exclude("Id",p.ID)

	//Prodlist = append(Prodlist, e)
}
func remove(s []*Product, i int) []*Product {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func getNextId() int {
	lp := Prodlist[len(Prodlist)-1]
	return lp.ID + 1
}

var Prodlist = []*Product{
	&Product{
		ID:   1,
		Name: "Tes",
	},
	&Product{
		ID:   2,
		Name: "Tes2",
	},
}
