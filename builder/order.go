package builder

import "fmt"

type drink []string
type stapleFood []string
type dish []string
type order struct {
	Drink      drink
	StapleFood stapleFood
	Dish       dish
}

func (o *order) print() {
	fmt.Println("酒水:", o.Drink)
	fmt.Println("主食:", o.StapleFood)
	fmt.Println("菜:", o.Dish)
}
