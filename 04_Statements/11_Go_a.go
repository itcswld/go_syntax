//goroutine --獨立併發執行緒
package stmt

import (
	"fmt"
	"time"
)

func GoStmt_a() {
	//run both code side by side
	go count("on") //"go" run this func in the  background and then continue to the next line immediately
	count("off")
}

/*
1 off
1 on
2 on
2 off
*/

func count(power string) {
	for i := 1; true; i++ {
		fmt.Println(i, power)
		time.Sleep(time.Millisecond * 500)
	}
}
