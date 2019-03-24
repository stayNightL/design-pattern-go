package led

import (
	"fmt"
	"testing"
)

func TestLed(t *testing.T) {
	led := Led{status: "off"}
	led.on()
	fmt.Println(led.status)
}
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
