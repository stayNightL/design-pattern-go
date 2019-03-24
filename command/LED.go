package led

import "fmt"

type Command interface {
	execute()
	undo()
}
type Led struct {
	status string
}

func (l *Led) on() {
	l.status = "ON"
}
func (l *Led) off() {
	l.status = "off"
}

type CommandImp struct {
	Led   *Led
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
