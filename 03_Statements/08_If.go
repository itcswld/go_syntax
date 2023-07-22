//control structures
package stmt

import (
	"fmt"
)

func IfStmt() {
	//if-lse
	fmt.Println("============= if-else =============")
	var isTrue *bool
	t := true
	isTrue = &t

	if isTrue == nil {
		fmt.Println("no value")
	} else if !*isTrue {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
	//reslut:true
}
