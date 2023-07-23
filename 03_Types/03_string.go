package Types

import (
	"fmt"
	"strings"
)

func StringTypes() {
	var name string
	name = "Eve"
	country := "Taiwan"
	fmt.Println(name, "is come from", country)
	fmt.Println(strings.Contains(name, country))    //false
	fmt.Println(strings.ReplaceAll(name, "e", "a")) //Eva

	//to array
	fmt.Println(strings.Split(name, ""))  //[E v e]
	fmt.Println(strings.Split(name, " ")) //[Eve]
	arrSplited := strings.Split(name, "")
	arrSplited1 := strings.Split(name, " ")
	fmt.Println(arrSplited[0])  //E
	fmt.Println(arrSplited1[0]) //Eve
}
