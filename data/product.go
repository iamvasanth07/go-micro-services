package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//Product Struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

//ToJSON used to convert data from product struct to json
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//FromJSON used to convert data from request to product struct
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//GetProduct returns all the product
func GetProduct() Products {
	return productList
}

//AddProduct add new product
func AddProduct(p *Product) {
	p.ID = getNextID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

//UpdateProduct updates the existing product
func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	p.UpdatedOn = time.Now().UTC().String()
	return nil
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

//ErrProductNotFound ...
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milky Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Frothy Milky Coffee",
		Price:       1.99,
		SKU:         "asd123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
