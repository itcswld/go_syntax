package Types

import "fmt"

///繼承

func (b *Bus) Run() {
	fmt.Println("bus is runing ")
}
func (b *Bus) Stop() {
	fmt.Println("bus stop")
}

type Engine interface {
	Run()
	Stop()
}

//繼承
type Bus struct {
	Engine
}

func (b *Bus) Working() {
	b.Run()
	b.Stop()
}

func InterfaceTypes_c() {

	fmt.Println("======== 繼承 ========")
	bus := Bus{}
	bus.Working()

}
