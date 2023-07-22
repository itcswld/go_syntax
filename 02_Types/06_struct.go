//object
package Types

import (
	"fmt"
	"reflect"
)

//struct
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

type SecretAgent struct {
	PPl
	license bool
}

func (sa SecretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "i wanna be a Go Dev. in 2022."`)
}

// func (s PPl) speak() {
// 	fmt.Printf("Hi, my name is %s %s.", s.fname, s.lname)
// }
