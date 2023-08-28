package Types

import (
	"fmt"
	"reflect"
)

func PointerTypes() {
	/*
		which include memory address"&", and value "*",
		it depens on memory address.
		 & obtain from (a pointer) the address of
			a data item held in another location.
	*/

	//PointerType = "*" BaseType .
	//set pointer to z int
	var z *int //像flutter int? z
	xLoc := z

	fmt.Println(z)  //nil, The value of an uninitialized pointer is nil.
	fmt.Println(*z) //invalid memory address or nil pointer dereference

	var x int
	xLoc = &x
	fmt.Println(x)     //0
	fmt.Println(xLoc)  //0x1400012c008
	fmt.Println(*xLoc) //0

	m1 := 2                                              //BaseType
	ptr := &m1                                           //& memory place(地址)
	fmt.Println(ptr)                                     //0x14000122008
	fmt.Println(reflect.TypeOf(ptr))                     //*int , PointerType (var m1 *int)
	fmt.Println(*ptr)                                    //2,"*"get value
	*ptr = 8                                             //change x value
	fmt.Println(*ptr, " //type :", reflect.TypeOf(*ptr)) //int

	//--Pointer type with struct
	type (
		Food struct{ fluit string }
	)
	var f *Food         //flutter: Food? f;
	fmt.Println(f)      //<nil>
	f = &Food{"Banana"} //init f
	fmt.Println(f.fluit)
}
