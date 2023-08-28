package Types

import (
	"fmt"
	"reflect"
)

func anytype(any interface{}) {
	fmt.Println(any, "'s type is :", reflect.TypeOf(any))
}

func InterfaceTypes_d() {

	anytype("string")                     //string 's type is : string
	anytype(1)                            //1 's type is : int
	anytype(struct{}{})                   //{} 's type is : struct {}
	anytype(struct{ name string }{"Eve"}) //{Eve} 's type is : struct { name string }

}
