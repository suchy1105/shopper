package shop

type Shop struct{
	Name 	string 
	Street 	string
	Number 	int
	City	string
	categoryStack[] string
}

func NewShop(name string, street string, number int, city string ) *Shop{
	s:=&Shop{
		Name: 	name
		Street:	street
		Number: number
		City:	city 
	}
	return s
}
func (s *Shop) Save() bool{
	bool success 
	//fill
	return success
}

func (s *Shop) Edit() bool{ 
	bool success 
	//fill
	return success
}