package builder

import (
	"log"
	"testing"
)

func TestOrder(t *testing.T) {
	order := order{
		Drink:      []string{"娃哈哈"},
		Dish:       []string{"宫保鸡丁"},
		StapleFood: []string{"面条"},
	}
	order.print()
}
func TestSimpbuilder(t *testing.T) {
	builder := SimpleBuilder{}
	b := builder.Init()
	var err error
	if b, err = builder.OrderDish("dish-宫保鸡丁"); err != nil {
		log.Fatal(err)
	}
	if b, err = b.OrderDrink("drink-娃哈哈"); err != nil {
		log.Fatal(err)
	}
	if b, err = b.OrderStapleFood("stapleFood-面条"); err != nil {

		log.Fatal(err)
	}
	o, _ := b.CreateOrder()
	o.print()
}
func TestSimpbuilder2(t *testing.T) {
	builder := SimpleBuilder{}
	b := builder.Init()
	var err error
	var o *order
	if b, err = b.OrderDrink("drink-娃哈哈"); err != nil {
		log.Fatal(err)
	}

	if o, err = b.CreateOrder(); err != nil {
		log.Fatal(err)

	}
	o.print()
}
