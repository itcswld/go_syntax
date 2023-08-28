//dictionary
package Types

import "fmt"

type (
	Prod struct {
		name  string
		price int
	}
)

func MapTypes() { //= "map" "[" Type "]" ElementType .
	//map[string]int
	var mi map[string]int
	mi = map[string]int{
		"Eve": 33,
		"Eva": 29,
	}
	fmt.Println(mi) //map[Eva:29 Eve:33]

	//map[*T]struct{ x, y float64 }
	var prod map[int]struct {
		name  string
		price int
	}
	prod = map[int]struct {
		name  string
		price int
	}{1: {"car", 1000}}
	fmt.Println(prod[1]) //{car 1000}

	//map[string]interface{}
	var m3 map[string]interface{}
	m3 = map[string]interface{}{"car": 1000}
	fmt.Println(m3["car"]) //1000

	//--- empty map ----
	//make(map[string]int)
	var m4 = make(map[string]string) //= m4:= map[string]string{}
	m4["last"] = "tong"
	name, isExist := m4["name"]
	fmt.Println("Is name Exist ?", isExist, ",", name) //true,eve
	last, isExist := m4["last"]
	fmt.Println("Is last name Exist ?", isExist, last) //false
	if lastName, isExist := m4["Last"]; isExist {
		fmt.Println("last name: ", lastName)
	} //; =then()

	//make(map[string]int, 5) --當容量超過上限5， 還是會自動加1， 所以5只是個注明
	m5 := make(map[int]int, 5)
	m5 = map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	fmt.Println(m5)

}
