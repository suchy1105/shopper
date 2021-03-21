package shoppinglist
import(
	"github.com/suchy1105/shopper"
)
//Shoppinglist object
type Shoppinglist struct {
	//	Date time.date
	productlist [] Product
	shop 		   shop.Shop
}

//NewShoppinglist create shopping list
func NewShoppinglist() *Shoppinglist {
	sl := &Shoppinglist{}
	return sl
}
func (s *Shoppinglist) sortbyshop() {

}
func (s *Shoppinglist) addProduct() {

}
func (s *Shoppinglist) removeProduct() {

}
