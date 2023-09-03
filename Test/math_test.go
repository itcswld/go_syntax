package test

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) { //func_name 必須是 Test開頭 ， ref 必須是 *testing.T
	arr := []int{3, 1, 5, 6}
	min := Min(arr)
	fmt.Println(min)
}

//go test
//go test -v
//go test -v -run="Test"

/*
=== RUN   TestMin
1
--- PASS: TestMin (0.00s)
PASS
ok      syntax/Test     0.104s
*/

//--生成執行檔
//go test -c
//go test -v -o test123.test
