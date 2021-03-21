package product

import (
	"fmt"
)

//Product producjt
type Product struct {
	Name        string
	Quantity    float32
	Category    string
	MeasureType string
}

//NewProduct returns new product pointer
func NewProduct(n string, q float32, c string, m string) *Product {
	p := &Product{
		Name:        n,
		Quantity:    q,
		Category:    c,
		MeasureType: m,
	}
	return p
}

//Delete delate
func Delete() {
	fmt.Println("Deleted")
}

//NotAvailable the
func NotAvailable() {
	fmt.Println("not available")
}
