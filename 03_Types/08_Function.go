package Types

import (
	"fmt"
)

func FuncType() {
	//FunctionType   = "func" Signature .
	func() {
		fmt.Println("good")
	}() //good

	func(x int) int {
		fmt.Println(x)
		return x
	}(1) //1

	func(a, _ int, z float32) bool {
		return a == int(z)
	}(1, 1, 0.1) //false

	func(s string, v ...int) {
		fmt.Println("s = ", s, " ; v =", v)
	}("good", 1, 2, 3) //s =  good ; v= [1 2 3]

	func(a, b int, z float64, opt ...interface{}) (success bool) {
		c := a + b
		fmt.Println(opt...)
		return c == int(z)
	}(1, 2, 3.0, 1, 2)

	func(int, int, float64) (float64, *[]int) {
		arrInt := []int{1, 2}
		return 3.0, &arrInt
	}(1, 1, 3.0)

	/*
		func(n int) func(p *int) {
			return func(c *int) {
				fmt.Println(*c)
			}
		}(1)
	*/
	funcA := closure()
	for i := 0; i < 10; i++ {
		fmt.Print(funcA(i, i+1), "\n")
	}

}

func closure() func(a, b int) int {
	return func(a, b int) int {
		fmt.Printf(" a * b = %d  * %d = ", a, b)
		return a * b
	}
}
