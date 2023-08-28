package Types

import (
	"fmt"
)

func closure() func(a, b int) int {
	return func(a, b int) int {
		fmt.Printf(" a * b = %d  * %d = ", a, b)
		return a * b
	}
}

func Min(arr [5]int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}

func GetElems(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func FuncType_a() {
	//FunctionType   = "func" ([Parameters]) [ Result(Parameters | Type ) ] {}

	//closure
	//func()
	func() {
		fmt.Println("good")
	}() //good

	//func(x int) int
	func(x int) int {
		fmt.Println(x)
		return x
	}(1) //1

	//func(a, _ int, z float32) bool | func(a, b int, z float32) (bool)
	func(a, _ int, z float32) bool {
		return a == int(z)
	}(1, 1, 0.1) //false

	//func(prefix string, values ...int) -- ...int參數不定量
	func(s string, i ...int) {
		fmt.Println("s = ", s, " ; i =", i)
	}("good", 1, 2, 3) //s =  good ; i= [1 2 3]

	//func(a, b int, z float64, opt ...interface{}) (success bool)
	fmt.Println()
	r0 := func(a, b int, z float64, opt ...interface{}) (Result bool) {
		c := a + b
		fmt.Println(opt...)
		Result = c == int(z)
		return
	}(1, 2, 3.0, 1, 2)
	fmt.Println("func(a, b int, z float64, opt ...interface{}) (success bool)", r0) //true
	//min
	num := [5]int{5, 3, 4, 2, 1}
	min := Min(num)
	fmt.Println(num, "; the min number is :", min) //1

	//func(int, int, float64) (float64, *[]int)
	r1, r2 := func(int, int, float64) (float64, *[]int) {
		arrInt := []int{1, 2}
		return 3.0, &arrInt
	}(1, 1, 3.0)
	fmt.Println("func(int, int, float64) (float64, *[]int)---", r1, r2) // 3 &[1 2]

	//func(n int) func(p *T)
	a := 1
	b := 2
	s := Swap{&a, &b}
	msg := fmt.Sprintf("before  swap a  = %d , b = %d", *s.a, *s.b)
	s.exchange(msg)

	funcA := closure()
	for i := 0; i < 10; i++ {
		fmt.Print(funcA(i, i+1), "\n")
	}

	//func(i []int], f func(int)) -- 用匿名函數（closure）作為 回呼函數
	list := []int{1, 2, 3}
	GetElems(list, func(v int) {
		fmt.Println(v)
	})

}

type (
	Swap struct {
		a, b *int
	}
)

//func(n int) func(p *T) --引用傳遞
func (s Swap) exchange(msg string) {
	s.a, s.b = s.b, s.a
	fmt.Println(msg, " after swap a = ", *s.a, ", b = ", *s.b)
}
