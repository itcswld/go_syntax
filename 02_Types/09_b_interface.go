//oop
package Types

import (
	"fmt"
)

//interface 物件導向 -> 共同的部份
type Person struct {
	fname string
	lname string
}
type Agent struct {
	PPl
	license bool
}
type human interface {
	speak()
}

func (p PPl) speak() {
	fmt.Println("person speaking!")
}
func (s Agent) speak() {
	fmt.Println("Agent speaking")
}
func whoIsSpeaking(h human) {
	h.speak()
}
func InterfaceTypes_b() {
	p := PPl{
		fname: "Eve",
		lname: "Tung",
	}
	a := Agent{
		p,
		true,
	}
	whoIsSpeaking(p)
	whoIsSpeaking(a)
}
