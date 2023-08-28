package Types

import "fmt"

//PointerType = "*" BaseType .
func swap(p1, p2 *string) { //交換值
	tmp := *p2
	*p2 = *p1
	*p1 = tmp
}

func exchange(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("*a type %T \n", *a)
	fmt.Printf("a  address: %p", a)
}

func PointerTypes_a() {
	//--pointer demenstrate swap 2 value
	p1, p2 := "Ben", "John"
	fmt.Println(p1, p2)
	swap(&p1, &p2) //&(地址)
	fmt.Println(p1, p2)

	a, b := 1, 2
	exchange(&a, &b)
	fmt.Println(a, b)
}
