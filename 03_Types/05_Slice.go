package Types

import (
	"fmt"
	"reflect"
)

//-- Slices are like \"references\" to arrays
func SliceTypes() {

	s := make([]int, 6, 10) //make([]T, length, capacity) , length must <= capacity, cap是設定預分配的元素數量， 可降低多次分配空間造成性能問題
	v := new([10]int)[:6]   //range 0~6
	//s[6] = 7  //index out of range [6] with length 6
	fmt.Println(s, v)                           //[0 0 0 0 0 0] [0 0 0 0 0 0]
	fmt.Println("s =", len(s), ", v =", len(v)) //s = 6 , v = 6

	num := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reflect.TypeOf(num)) //SliceType = "[" "]" ElementType .
	var slice []int = num[1:5]
	fmt.Println(slice)
	//Slices are like references to arrays
	s1 := num[0:2]
	s2 := num[1:3]              //range [1]~[3-1]; alike swift 1...5 not inculde 5
	s3 := num[:len(num)-2]      //from first one
	s4 := num[1:]               //to last one
	s5 := num[:]                //all
	fmt.Println(s1, s2, s3, s4) //[1 2] [2 3] [1 2 3 4] [2 3 4 5 6]
	s1[0] = 2
	fmt.Println(s1, s5) //[2 2] [2 2 3 4 5 6]
}
