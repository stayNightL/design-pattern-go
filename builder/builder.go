package builder

import (
	"errors"
	"strings"
)

// Builder 建造者接口
type Builder interface {
	CreateOrder() (*order, error)
	OrderDish(string) (Builder, error)
	OrderDrink(string) (Builder, error)
	OrderStapleFood(string) (Builder, error)
}

//SimpleBuilder 建造者简单实现
type SimpleBuilder struct {
	dish       dish
	drink      drink
	stapleFood stapleFood
}

// Init initialization
func (sb *SimpleBuilder) Init() Builder {
	sb.dish = dish{}
	sb.drink = drink{}
	sb.stapleFood = stapleFood{}
	return sb
}

// OrderDish add dish
func (sb *SimpleBuilder) OrderDish(dish string) (Builder, error) {
	if strings.HasPrefix(dish, "dish-") {
		sb.dish = append(sb.dish, strings.Split(dish, "-")[1])
		return sb, nil
	}
	return nil, errors.New("its not right type")
}

// OrderDrink add Drink
func (sb *SimpleBuilder) OrderDrink(drink string) (Builder, error) {
	if strings.HasPrefix(drink, "drink-") {
		sb.drink = append(sb.drink, strings.Split(drink, "-")[1])
		return sb, nil
	}
	return nil, errors.New("its not right type")
}

// OrderStapleFood add food
func (sb *SimpleBuilder) OrderStapleFood(stapleFood string) (Builder, error) {
	if strings.HasPrefix(stapleFood, "stapleFood-") {
		sb.stapleFood = append(sb.stapleFood, strings.Split(stapleFood, "-")[1])
		return sb, nil
	}
	return nil, errors.New("its not right type")
}

//CreateOrder create order
//*order sd
func (sb *SimpleBuilder) CreateOrder() (*order, error) {
	if len(sb.dish) == 0 && len(sb.stapleFood) == 0 {
		if len(sb.drink) == 0 {
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
