package product

import(
	"fmt"
)
type product struct{ 
 	Name string
 	Quantity float
	Category string
	MeasureType string
}
func NewProduct(n string, q float, c string, m MeasureType)*product{
	p:= &product{
		Name:		n,
		Quantity:	q,
		Category:	c,
		MeasureType	m,
	}
	return p
}
func Delete(){
	fmt.Println("Deleted")
}
func NotAvailable(){
	fmt.Println("not available")
}
