package builder

import "errors"

type Builder interface {
	CreateOrder() (*order, error)
	orderDish(string) (*Builder, error)
	orderDrink(string) (*Builder, error)
	orderStapleFood(string) (*Builder, error)
}
type SimpleBuilder struct {
	dish       dish
	drink      drink
	stapleFood stapleFood
}

func (sb *SimpleBuilder) orderDish(dish string) *Builder {
	
}
func (sb *SimpleBuilder) orderDrink(drink string) *Builder {

}
func (sb *SimpleBuilder) orderStapleFood(stapleFood string) *Builder {

}
func (sb *SimpleBuilder) CreateOrder() (*order, error) {
	if len(sb.dish) == 0 && len(sb.stapleFood) == 0 {
		if len(sb.dish) == 0 {
			return nil, errors.New("请重新点餐")
		}
		return nil, errors.New("饮料不能单点")

	}
	order := order{
		Dish:       sb.dish,
		Drink:      sb.drink,
		StapleFood: sb.stapleFood,
	}
	return &order, nil

}
