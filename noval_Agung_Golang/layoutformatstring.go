package novalagunggolang

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isgraduated bool
	hobbies     []string
}

var data = student{
	name:        "sukma",
	height:      155.5,
	age:         25,
	isgraduated: false,
	hobbies:     []string{"eat", "sleep"},
}

// layout format %b
func mainExample() {
	fmt.Printf("%b\n", data.age)

	/*Layout Format %c digunakan untuk memformat data numerik yang merupakan code unicode, menjadi bentuk string karakter unicodenya
	 */
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)
	/*Layout format %d
	Digunakan untuk memformat data numerik, menjadi string numerik berbasis 10 (basis bilangan yang kita gunakan )*/
	fmt.Printf("%d\n", data.age)

}
