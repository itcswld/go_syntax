package pack

import (
	"fmt"
	"reflect"
)

//2.反射類型物件 轉 介面類型變數
//type & value打包存到一個介面類型變數 func (v value) Interface() interface{}
func ReflectPack_b() {

	var n interface{} = "Eve"
	fmt.Printf("orign type = %T , value = %v \n", n, n) //orign type = string , value = Eve

	//2.反射類型物件 轉 介面類型變數
	t := reflect.TypeOf(n)
	v := reflect.ValueOf(n)
	fmt.Printf("Reflect's Type = %T , Value's Type = %T \n", t, v) //Reflect Type = *reflect.rtype , Value = reflect.Value

	//1.介面類型轉反射類型物件
	i := v.Interface()
	fmt.Printf("v type = %T , value = %v \n", i, i) //v type = string , value = Eve

}
