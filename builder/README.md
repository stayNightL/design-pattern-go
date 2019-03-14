### 建造者模式
#### 场景
考虑一个餐馆的点餐系统，该系统需要在顾客点餐之后提供点餐信息和打印小票，餐品分为酒水饮料，主食，菜品三类。

#### 代码V1.0
简单思考之后，需要一个代表订单的类，该类有一个负责打印的方法，和三个属性分别代表三类餐品。具体代码如下：
```

type drink []string
type stapleFood []string
type dish []string
type order struct {
    Drink drink
    StapleFood stapleFood
    Dish dish
}
func (o *order) print() {
    fmt.Println("酒水:", o.Drink)
    fmt.Println("主食:", o.StapleFood)
    fmt.Println("菜:", o.Dish)
}


```
测试代码：
```
func TestOrder(t *testing.T) {
    order := order{
        Drink: []string{"娃哈哈"},
        Dish: []string{"宫保鸡丁"},
        StapleFood: []string{"面条"},
    }
    order.print()
}
```
打印：
```
酒水: [娃哈哈]
主食: [面条]
菜: [宫保鸡丁]
```
#### 引入变化
首先思考这样一个问题：
1. 在测试用例里面，人为保证了娃哈哈是饮料，面条是主食，但是并没有在程序里面限制--如果硬要把宫保鸡丁放在饮料里面，也是可以正常运行的。
如何限制这种行为呢？
2. 如果饮料不可以单点但是可以先点呢？
3. 如果点三个菜以上送一瓶饮料呢？
4. 如果餐馆打算推出套餐呢？
#### 分析变化
对于1：可以在单品名称上引入类型标识辅助判断，方法可以绑定在order，但是会有问题。
对于2：延迟order对象的创建--因为只有点完餐了才知道饮料是否是单点的
对于3：解决方法同2，延迟对象创建之后可以很方便的做这件事。
对于4：把套餐内容定义在辅助创建对象的具体实现子类即可。
#### 建造者模式--隐藏和影响对象创建过程
1. 首先引入builder接口：
```
type Builder interface {
	CreateOrder() (*order, error)
	OrderDish(string) (Builder, error)
	OrderDrink(string) (Builder, error)
	OrderStapleFood(string) (Builder, error)
}
```
2. 满足1，2，3的一个简单实现：
```
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
```
测试代码：
```
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

```
#### 思考
如果你在之前了解过建造者模式，你可能会觉得本文会少一个很重要的部分，builder类的所属者--真正执行builder的对象。是的，本文的确缺少这一部分。因为这不是想展现的主要部分，而且测试代码也在一定程度上表现了这一部分--如何去使用builder。
建造者模式的效果有以下三个：
1. 它使你可以改变一个产品的内部表示
2. 它将构造代码和表示代码分开
3. 它使你可对构造过程进行更精细的控制
本文展现的最主要部分是3，首先不允许错误的点餐，然后是饮料不能单点，满足不了上述任一一条都无法构造出正确的order对象，这些都是对构造过程的控制。

ps：这个案例真的适合表达建造者模式吗？欢迎提issue。
