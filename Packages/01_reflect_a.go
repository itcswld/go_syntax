//reflection
package pack

import (
	"fmt"
	"reflect"
)

//1.介面類型轉反射類型物件
func ReflectPack_a() {
	var v1 int = 1
	fmt.Println("v init value: ", v1)
	fmt.Println("v type: ", reflect.TypeOf(v1))
	v2 := reflect.ValueOf(v1)
	fmt.Println(v1)                                                           //2
	fmt.Println(v2.Type())                                                    //int
	fmt.Println(v2.Kind())                                                    //int
	fmt.Println(v2.Int())                                                     //1
	fmt.Println("n & n_v1 are same type ? ", reflect.TypeOf(v1) == v2.Type()) //true
}
