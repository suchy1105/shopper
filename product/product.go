package product

import (
	"fmt"
)

//Product producjt
type Product struct {
	Name        string
	Quantity    float32
	Category    string
	Subcat		string
	Type		string
	MeasureType string
	ID 			int
	Mass		int //grams
	PiecePrice	float32
	Mark		string
}

//NewProduct returns new product pointer
func NewProduct(n string, q float32, c string, sc string, t string ,m string, id int, ms int, pp float32, tm string) *Product {
	p := &Product{
		Name:        n,
		Quantity:    q,
		Category:    c,
		Subcat: 	 sc,
		Type: 		 t,
		MeasureType: m,
		ID: 		 id,
		Mass:  		 ms,
		PiecePrice:  pp,
		Mark:		 tm,
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
