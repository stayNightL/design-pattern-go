### 命令模式
#### 一个可任意调节的LED灯
假设有这么一款LED灯，有开、关两个个操作
代码如下：
```
type Led struct {
    status string
}
func (l *Led) on(){
    l.status="ON"
}
func (l *Led) off(){
    l.status="off"
}
```
现在要在此基础上引入撤销操作。这里我们虽然可以简单的添加一个undo方法，通过简单的逻辑判断来将off转换为on或者on转换为off，但是这只是一个假撤销，当状态转换为off->off->off的时候，实际上是不对的，off撤销完应该还为off。对于此，我们需要把对Led的每次调用抽象为一次命令，命令内部封装了撤销方法和调用处的‘现场状态’，对Led方法的调用转换为对一系列命令对象的调用，调用处会维护一个命令队列，这样撤销时只需要出栈并调用撤销方法即可。
#### 代码实现
```

type CommandImp struct {
    Led *Led
    staus string
}
func (oc *CommandImp) execute() {
    fmt.Println("no such execute!")
}
func (oc *CommandImp) undo() {
    if oc.staus == "off" {
        oc.Led.off()
    } else {
        oc.Led.on()
    }
    fmt.Println("status: ", oc.staus, "->", oc.Led.status)
}
type onCommand struct {
    CommandImp
}
func (oc *onCommand) New(led *Led) {
    oc.Led = led
    oc.staus = oc.Led.status
}
func (oc *onCommand) execute() {
    oc.staus = oc.Led.status
    oc.Led.on()
    fmt.Println("status: ", oc.staus, "->", oc.Led.status)
}
type offCommand struct {
    CommandImp
}
func (oc *offCommand) execute() {
    oc.staus = oc.Led.status
    oc.Led.off()
    fmt.Println("status: ", oc.staus, "->", oc.Led.status)
}
func (oc *offCommand) New(led *Led) {
    oc.Led = led
    oc.staus = oc.Led.status
}


```
代码测试调用：
```

func TestUndo(t *testing.T) {
    his := []Command{}
    led := &Led{status: "off"}
    on := onCommand{}
    on.New(led)
    off := offCommand{}
    off.New(led)
    on.execute()
    his = append(his, &on)
    off.execute()
    his = append(his, &off)
    //逆向循环模仿出栈
    for i := len(his) - 1; i >= 0; i-- {
        his[i].undo()
    }
    fmt.Println(led.status)
}


```
#### 重要的思考
在代码中我们做了两件事：
1. 把方法调用封装成命令
2. 每种命令有自己的对应的逆操作用来撤销--虽然这个例子里的撤销逻辑是一样的。
但是如果你仔细看测试代码，会发现更有意思的事情，主要表现在这几行：
```
on.execute()     
his = append(his, &on)     
off.execute()     
his = append(his, &off)     
//逆向循环模仿出栈     
for i := len(his) - 1; i >= 0; i-- {         
his[i].undo()     
}

```
这是我们创建对象之后的代码，有趣的是：
**在执行或者撤销的过程中，没有任何关于Led对象的信息暴露出来！**
这意味着命令模式可以把 **执行逻辑和具体对象分离，对象之间不需要有任何关系**

#### 以用命令模式来实现的功能例子还有：

* 交易行为

* 进度列
* 向导
* 用户界面按钮及功能表项目
* 线程 pool
* 宏收录