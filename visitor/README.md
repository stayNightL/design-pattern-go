### 访问者模式
#### 从订单开始
假设我们手里有这么一个订单，每一种商品只能添加一个，我们需要计算出所有商品一共的价格，那么，第一版可能是这样的：

```
type Order struct {
	goods map[string]float32 // key :物品，value：价格,假定一件东西只能买一件
}
type OrderInterface interface {
	order() float32
}

func (o *Order) order() float32 {
	var count float32
	for _, v := range o.goods {
		count += v
	}
	return count
}

```
调用过程如下：
```

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
```
#### 引入变化
1. 如果满1000-10这么办？
2. 如果买三免一怎么办？
3. 如果三件八折怎么办？
作为一个有经验的程序员，当然可以很快给出第二版：在Order身上再添加上1，2，3 对应的方法呗！
这个思路当然行得通，但是如果活动再加上下面几个呢？
3. 四件7折
4. 买五免二
如果3和4 同时生效的话。。。。
Order的代码会爆炸，调用处也会充满奇奇怪怪的判断!
更重要的是，每次都需要修改order的代码！严重违反了对修改关闭，对扩展开放的原则。
#### 分析变化
按照引入变化的内容，可以很清晰的知道，变化的只是商品总价的计算方法，不变的是商品订单的结构--订单永远是个map。
#### 访问者模式--分离数据结构和算法
1. 把Order中的goods 和 order方法分离，goods仍然在Order中，order方法放到另一个名为visitor的接口中作为接口方法（这个名称并不合适，谁让访问者是主角呢）。
2. 在Order中新增一个方法：accept,该方法接收一个visitor，方法中使用外部提供的算法（也就是order方法），并为其提供数据结构，，进行计算和返回。

代码如下：
```
type Order struct {
	goods map[string]float32 // key :物品，value：价格,假定一件东西只能买一件
}
type OrderInterface interface {
	accept(visitor *visitor) float32
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
```
这里只实现了一个商品八折的算法。
#### 思考
在访问者模式中，我们做的事情其实很少：只是把一个方法移到另一个抽象接口中，让数据和逻辑（数据结构和算法）分离，这样一来，添加逻辑或者算法都不需要对数据结构进行修改，而是新增接口的子类。

就这个例子来说，单纯的访问者模式的应用并没有什么及其吸引人的地方--因为它除了分离并没有什么其他作用，但是如果它跟其他模式结合起来呢？比如建造者模式，是不是就可以动态的修改算法的细节?
这可是初版的代码做不到的事情。但是建造者模式不在此次话题之内，有兴趣的可以自己动手。