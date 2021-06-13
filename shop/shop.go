package shop

import (
	//"github.com/suchy1105/shopper/api"
	"github.com/suchy1105/shopper/product"
)

//Shop strict
type Shop struct {
	Name          string
	Street        string
	Number        int
	City          string
	Products	  []product.Product
	CategoryStack []string
	ID 			  int
	WidGPS		  float32
	HigGPS		  float32
}

//NewShop create new shop
func NewShop(name string, street string, number int, city string, id int, wgps float32, hgps float32) *Shop {
	s := &Shop{
		Name:   name,
		Street: street,
		Number: number,
		City:   city,
		ID: id,
		WidGPS: wgps,
		HigGPS: hgps,
	}

	return s
}

//Save end modification
func (s *Shop) Save() bool {
	var success bool
//	api.Push()
	return success
}

//Edit edit shop
func (s *Shop) Edit(name string, street string, number int, city string, id int, wgps float32, hgps float32) bool {
	var success bool
	s.Name=name
	s.Street=street
	
	return success
}
