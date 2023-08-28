//oop
package Types

import (
	"fmt"
)

//interface 物件導向
/*
oop 物件導向的根類型
	1.封裝:
		1.屬性
		2.方法
		3.權限 //其他pacage, 大寫=public, 可存取； 小寫=private，不可存取
	2.繼承
	3.多形 -> 導向共同的部份
*/

//封裝-屬性
type Student struct {
	name string
}

func (s *Student) GetName() string {
	return s.name
}

//封裝-方法
func (s *Student) SetName(newName string) {
	s.name = newName
}

//封裝-屬性
type Cat struct {
	name string
}
type Dog struct {
	name string
}

//多形 -- 根據物件定義實現不同行為 =  實現多形行為
//make common function as interface ,call same func  do (eat)  different things
type Action interface {
	eat()
}

func act(a Action) {
	a.eat()
}

//封裝-方法
func (c Cat) eat() {
	fmt.Println(c.name, " eat fish. ")
}
func (d *Dog) eat() {
	fmt.Println(d.name, "eat meat.")
}

func InterfaceTypes_b() {

	fmt.Println("======== 方法 ========")
	//獲取和設定屬性
	s := new(Student)
	s.SetName("Eve")
	fmt.Println(s.GetName())

	fmt.Println("======== 多形 ========")
	c := Cat{"Kitty"}
	d := Dog{"Snoopy"}
	act(c)  //Kitty  eat fish.
	act(&d) //Snoopy eat meat.
}
