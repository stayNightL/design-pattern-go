package visitor

type Order struct {
	goods map[string]float32 // key :物品，value：价格,假定一件东西只能买一件
}
type OrderInterface interface {
	accept(visitor *visitor) float32
}

func (o *Order) order() float32 {
	var count float32
	for _, v := range o.goods {
		count += v
	}
	return count
}
func (o *Order) accept(visitor visitor) float32 {
	return visitor.order(o)
}

type visitor interface {
	order(o *Order) float32
}
type off20Visitor struct {
}

func (off20 *off20Visitor) order(o *Order) float32 {
	var count float32
	for _, v := range o.goods {
		count += v
	}
	return count * 0.8
}
