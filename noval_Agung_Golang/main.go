package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	var secret interface{}

	secret = "Sukma Rizki"
	fmt.Println(secret)

	secret = map[string]int{"sukma": 1, "rizki": 2}
	fmt.Println(secret)
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "Sukma",
		"age":       23,
		"breakFast": []string{"apple", "banana", "watermelon"},
	}
	fmt.Println(data)
	//casting variabel interface kosong ke objek Pointer
	type person struct {
		name string
		age  int
	}

	var empty interface{} = &person{name: "Sukma", age: 28}
	var name = empty.(*person).name
	var age = empty.(*person).age
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("Sukma Rizki")

	/*Akses Information Variabel Objek */
	var s1 = &student{name: "john", Grade: 3}
	fmt.Println("nama :", s1.name)

	var reflectvalue = reflect.ValueOf(s1)
	var method = reflectvalue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})
	fmt.Println("nama: ", s1.name)
	runtime.GOMAXPROCS(2)
	go print(5, "halo")
	print(5, "Apa kabar")

	var input string
	fmt.Scanln(&input)

	//Menggunakan fungsi fmt.scanln()
	/*Fungsi ini mengcapture semua karakter sebelum user menekan tombol enter 	
	lalu menyimpannya pada variabel kode diatas merupakan skema fungsi fmt.scanln fungsi tesebut 
	bisa menampung tipe data bertipe interface berjumlah tak terbatas.tiap parameter akan menampung karakter karakter 
	inputan user yang sudah dipisahkan dengan tanda spasi agar lebih jelas silahkan perhatikan contoh berikut 
	*/
	var  s2, s3, s4 string
	fmt.Scanln(&s2,&s3,&s4)


	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	//Penerapan chanel berbentuk variabel, dibuat dengan kombinasi keyword make dan chan 
	/*Program berikut adalah contoh implementasi chanel. 3buah goroutine dieksekusi, dimasing-masing goroutine
	
	terdapat proses pengiriman lewat chanel. kesemua data tersebut nantinya diterima oleh goroutine utama yaitu proses 
	proses yang dijalankan ddidalam fungsi blok main().
	*/
}

type student struct {
	name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.name = name
}

//Goroutine secara konsep mirip seperti thred, meskipun sebenarnya berbeda Sebuah native thread bisa
// bisa berisikan banyak goroutine mungkin lebih pas kalau gorotuine disebut mini thread Groutine
// sangat ringan hanya membutuhkan 2 kb saja untuk satu buah goroutine.Eksekusi goroutine bersifat asyncronous, menjadikanya tidak
// saling tunggu dengan goroutine lain
/*Goroutine merupakan salah satu paling penting dalam pemograman concurrenct.Salah satu membuat goroutine sangat istimewa adalah eksekusinya dijalankan
di multicoreproccessor kita bisa banyak temukan berapa core yang aktif makin banyak makin cepat

Penerapan Gouroutine proses yang dibutuhkan dalam gorutine proses yang akan dieksekusi
Brikut merupakan contoh implementasi sederhana tentang goroutine program dibawah ini menampilkan 10 baris teks , 5 eksekusi dengan cara biasa dan 5 lainya dieksekusi sebagai
goroutine baru
*/
func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
