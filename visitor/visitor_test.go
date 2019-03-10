package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	order := Order{
		goods: make(map[string]float32),
	}
	//添加商品
	order.goods["固态硬盘"] = 2333
	order.goods["16G内存"] = 1444
	order.goods["2080TI"] = 4396
	order.goods["i7-8"] = 1900
	//计算价格
	fmt.Println(order.order())

}

func TestVisitor2(t *testing.T) {
	order := Order{
		goods: make(map[string]float32),
	}
	//添加商品
	order.goods["固态硬盘"] = 2333
	order.goods["16G内存"] = 1444
	order.goods["2080TI"] = 4396
	order.goods["i7-8"] = 1900
	//计算价格
	visitor := off20Visitor{}

	fmt.Println(order.accept(&visitor))

}
