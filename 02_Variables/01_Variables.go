package Variables

import (
	"fmt"
)

func Variables() {

	//--var
	var age int // age is nil and has static type int
	age = 33
	fmt.Println(age)

	//mutiple variables
	var n1, n2 int = 1, 2
	var (
		year = 77
		bc   = 1911
	)
	var n4, n5 = 5, "hi"
	fmt.Println(n4, n5)

	//initializers :
	month := 2
	month = n1 + n2 // month += 1
	n1, n2, n3 := 5, 5, 4
	var date int = n1 + n2 + n3
	fmt.Println("i am ", age, " years old. bron in ", year+bc, "/", month, "/", date)

}
