package builder

import "testing"

func TestOrder(t *testing.T) {
	order := order{
		Drink:      []string{"娃哈哈"},
		Dish:       []string{"宫保鸡丁"},
		StapleFood: []string{"面条"},
	}
	order.print()
}
