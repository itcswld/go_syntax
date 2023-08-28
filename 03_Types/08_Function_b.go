package Types

import "fmt"

//defer --“反序”或”延遲敘述“， 為了在函數執行完後即時釋放資源。
/*
(1)關閉資源：釋放記憶體
（2）和recover()函數一起用：當程式當機或panic時，recover()可恢復執行，不會報當機錯誤，因在當機前執行defer， 依次來恢復程式
	return 或 panic 異常導致函數結束之前最後執行defer
*/
func FuncType_b() {
	//defer 會儲存程序疊依序執行，先進最後執行
	normalFunc("1st func")                 //1
	defer deferFunc("1st defer")           //5
	fmt.Println(returnFunc("Return func")) //2
	defer deferFunc("2sec defer")          //4
	normalFunc("2sec func")                //3

	/*
		1st func
		        defer from 1st func
		defer from Return func
		Return func
		2sec func
		        defer from 2sec func
		2sec defer
		1st defer
	*/
}

func deferFunc(s string) {
	fmt.Println(s)
}

func returnFunc(s string) (r string) {
	defer fmt.Println("defer from", s)
	return s
}

func normalFunc(s string) {
	defer fmt.Println("	defer from", s)
	fmt.Println(s)
}
