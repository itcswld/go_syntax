package Types

import (
	"fmt"
)

func ArrTypes() {
	//ArrayType   = "[" ArrayLength "]" ElementType .
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	arrInt := []int{1, 2, 3, 4, 5, 6}
	arrString := []string{"hello"}
	arrString = append(arrString, "World")
	fmt.Println(arr)
	fmt.Println(len(arr)) //length
	fmt.Println(arrInt, arrString)
}
