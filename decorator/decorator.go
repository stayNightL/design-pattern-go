package decorator

import "fmt"

type printInterface interface {
	print(string)
}
type print struct {
}

func (*print) print(content string) {
	fmt.Println(content)
}

type myprint struct {
	p printInterface
}

func (*myprint) print(content string) {
	fmt.Println("********")
	fmt.Println(content)
	fmt.Println("********")
}
