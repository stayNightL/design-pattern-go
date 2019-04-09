## 装饰者模式

> 装饰者模式是动态（运行期）向对象添加/强化功能的一种设计模式，该模式从类结构的角度上讲，装饰者和功能子类是同一级，但是其实装饰者是对功能子类的一种包装。也就是说装饰类和功能子类是has-a的关系

### 打印机

首先我们先实现一版打印机，该类十分简单，只有一个print方法。
```go
type print struct {
}

func (*print) print(content string) {
	fmt.Println(content)
}
```

### 2.0--如何提供打印页眉、页脚功能
对于上面的1.0版本来说，直接修改也不失为一个便捷办法，但是，从面向对象的角度来说：

 1. 一个类同时提供了打印内容和样式的功能，违反单一职责原则
 2. 从使用的角度来说，不可能全部都需要打印页眉和页脚
 3. 违反对修改关闭，对扩展开放原则

综上所述，我们要做的就是：

> 在不改变现有类结构的情况下，添加页眉页脚的功能

答案自然是本文的主角：**装饰者模式**

### 2.0 代码

```go
type printInterface interface {
	print(string)
}
type myprint struct {
	p printInterface
}

func (*myprint) print(content string) {
	fmt.Println("********")
	fmt.Println(content)
	fmt.Println("********")
}
```
在上面的代码中，先抽取了打印功能的接口，其子类myprint包含了一个接口printInterface成员变量用以接收print的实例，测试代码如下：
```go
func TestDe2(t *testing.T) {
	p := myprint{p: &print{}}
	p.print("hello，decorator")
}
```
打印：
```
  ********
  hello，decorator
  ********
```
> 最重要的东西--装饰者对功能的扩展（组合）开放
### 适用场景
1. 对于一个模块可以分解成几种功能的任意组合的实现，但是又不确定到底需要哪几种组合或者是运行时才能确定功能组合方式
2. 安全的扩展一个类，同时没有副作用