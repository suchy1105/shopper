package shoppinglist
import(
	"github.com/suchy1105/shopper/shop"
	"github.com/suchy1105/shopper/product"
)
//Shoppinglist object
type Shoppinglist struct {
	//	Date time.date
	productlist [] product.Product
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
