package Types

import (
	"fmt"
	"math"
)

func NumericTypes() {
	//--int32 vs int64
	fmt.Println("--- int32 vs int64 ---")
	var n1 int32
	var n2 int64
	fmt.Println(int64(n1) + n2) //0 , because default value is 0

	//無明確類型的常數類型 ex math.Pi
	var a1 float32 = math.Pi
	var b1 complex128 = math.Pi
	//vs 被確定類型的常數
	const pi64 float64 = math.Pi
	var a2 float32 = float32(pi64)
	var b2 complex128 = complex128(pi64)
	fmt.Println("float32: a1=", a1, "vs a2=", a2)
	fmt.Println("complex128: b1=", b1, "vs b2=", b2)
}
