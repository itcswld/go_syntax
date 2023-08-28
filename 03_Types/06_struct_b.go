package Types

import "fmt"

type (
	Book struct{ name, auther string }
)

func printBook(b *Book) {
	fmt.Printf("this is book of %s \n", b.name)
	fmt.Printf("auther: %s", b.auther)
}

func Struct_Pointer() {
	book1 := Book{"Go", "google"}
	printBook(&book1)
}
