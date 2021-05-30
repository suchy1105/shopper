package shop

import (
	"github.com/suchy1105/shopper/api"
)

//Shop strict
type Shop struct {
	Name          string
	Street        string
	Number        int
	City          string
	categoryStack []string
}

//NewShop create new shop
func NewShop(name string, street string, number int, city string) *Shop {
	s := &Shop{
		Name:   name,
		Street: street,
		Number: number,
		City:   city,
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
func (s *Shop) Edit() bool {
	var success bool
	//fill
	return success
}
