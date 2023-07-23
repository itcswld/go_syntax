//object
package Types

import (
	"fmt"
	"reflect"
)

//StructType    = "struct" "{" { FieldDecl ";" } "}" . FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
//IdentifierList Type
type Stud struct {
	no   int
	name string
}

func (s1 Stud) getName() string { //return string type value
	return s1.name
}

func (s Stud) speak() {
	fmt.Println("Student no.", s.no, s.name, `says, "Goood morning, Eve"`)
}

//
type (
	T5 struct{ f *T5 }       // T5 contains T5 as component of a pointer
	T6 struct{ f func() T6 } // T6 contains T6 as component of a function type
	T7 struct{ f [10][]T7 }  // T7 contains T7 as component of a slice in an array
)

func StructTypes() {
	//struct
	s1 := Stud{1, "Chi"}
	s2 := Stud{
		no:   2,
		name: "Eve",
	}

	fmt.Println(struct{ class, name string }{"computer science", "eve"}) //{computer science eve}
	fmt.Println(s1, s2)                                                  //{1 Chi} {2 Eve}
	fmt.Println(s1.getName())                                            //Chi
	fmt.Println(s2.getName())                                            //Eve
	fmt.Println(reflect.TypeOf(s1))                                      //main.Student
	s1.speak()                                                           //Student no. 1 Chi says, "Goood morning, Eve"

	//struct
	sal := SecretAgent{
		PPl{
			"Eve",
			"Tung",
		},
		true,
	}

	sal.speak()     //Eve Tung says, "i wanna be a Go Dev. in 2022."
	sal.PPl.speak() //Hi, my name is Eve Tung.%

}

type PPl struct {
	fname string
	lname string
}

//FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
//EmbeddedField
type SecretAgent struct {
	PPl     //EmbeddedField
	license bool
}

func (sa SecretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "i wanna be a Go Dev. in 2022."`)
}
