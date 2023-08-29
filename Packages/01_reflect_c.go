package pack

import (
	"fmt"
	"reflect"
)

//3.修改【反射類型物件】， 其值必須可「可寫入」
func ReflectPack_c() {
	var n int = 1
	v1 := reflect.ValueOf(n)
	fmt.Println("Can v be changed ? ", v1.CanSet()) //Can v be changed ?  false
	//v1.SetInt(1)                          //panic: reflect: reflect.Value.SetString using unaddressable value
	v2 := reflect.ValueOf(&n)
	fmt.Println("Can v2 be changed ? ", v2.CanSet())
	//修改入口
	v3 := v2.Elem()
	v3.SetInt(2)
	fmt.Printf("after change Type = %T, Value = %v ", v3, v3)

}
