package Types

import (
	"fmt"
	"reflect"
)

func SliceTypes() {

	s := make([]int, 50, 100) //make([]T, length, capacity) , 100 ttl length, the length must <= capacity
	v := new([100]int)[0:50]  //range 0~50
	fmt.Println("s =", len(s), "= v =", len(v))

	arrInt := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reflect.TypeOf(arrInt))
	//SliceType = "[" "]" ElementType .
	var slice []int = arrInt[1:5]
	fmt.Println(slice) //2,5 #range [1]~[5-1]; alike swift 1...5 not inculde 5
	//Slices are like references to arrays
	fmt.Println("-- Slices are like \"references\" to arrays --")
	nameArr := []string{"Chi", "Eve", "Eva", "Eric"}
	fmt.Println(nameArr)
	nameSlice1 := nameArr[0:2] //Chi, Eve
	nameSlice2 := nameArr[1:3] //Eve, Eva
	fmt.Println(nameSlice1, nameSlice2)
	nameSlice1[1] = "Siang"
	fmt.Println(nameSlice1, nameSlice2, nameArr)
}
