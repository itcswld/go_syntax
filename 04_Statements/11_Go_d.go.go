//goroutine
package stmt

import (
	"fmt"
	"time"
)

//in go when the main goroutine finishes the program exits. regardness of what any other goroutines might be doing.
func GoStmt_d() {
	go counter("sleep") //"go" run this func in the  background and then continue to the next line immediately
	go counter("fish")  //continue immediately to this line.
	//for fix "go" havent had time to do anything.
	fmt.Scanln()
}

func counter(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
