package decorator

import "testing"

func TestDe(t *testing.T) {
	p := print{}
	p.print("hello，decorator")
}
func TestDe2(t *testing.T) {
	p := myprint{p: &print{}}
	p.print("hello，decorator")
}
