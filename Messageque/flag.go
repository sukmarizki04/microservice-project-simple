package messageque

import (
	"flag"
	"fmt"
)

//Penggunaan Flag
/*Flag memiliki kegunaan yang sama seperti argument, yaitu
untuk parameterize eksekusi program dengan penulisan dalam bentuk key dan value. Berikut merupakan contoh
penerapanya*/

func mainFlag(){
	var name = flag.String("name", "anonymous", "typeyourname")
	var age = flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name \t : %s \n", *name)
	fmt.Printf("age\t: %d\n", *age)
}

//cara penulisan argument menggunakan flas
// go run bab45.go -name="john wick" -age=28
/*Tiap argument harus ditentukan key, tipe data dan nilai defaultnya.Contohnya 
seperti pada flag.String() diatas agar lebih mudah dipahami mari kita bahas kode berikjt

*/
var dataName = flag.String("name", "anonymous", "type yourname")
fmt.Println(*dataName)
/*Kode tersebut maksudnya adalah disiapkan flag betipe string 
dengan key adalah name dengan nilai default anonymous dan keterangan "type your name", Nilai flanya sendiri akan disimpan kedalam variabel

Nilai balik fungsi slag.string adalah pointer jadi perlu di-deference 
terlebih dahulu untuk mengakses nilai aslinya (*dataName)
Flag yang nilainya tidak di st secara otomatis akan mengembalikan nilai default 
Tabel berikut merupakan macam-macam fungsi flag yang tersedia tiap jenis data 
tipe data 

Deklarasi Flag dengan cara Passing Reference Variabel Penampung Data 
Sebenarnya ada dua cara deklarasi flaf yang bisa digunakan dan cara diatas merupakan cara pertama
cara kedua mirip dengan cara pertama perbedaanya kalau dicara pertma nilai pointer flag dikembalikan lalu ditampung variabel sdangkan pada cara kedua, nilainya diambil lewat 
parameter 
*/

//cara pertama
var data1 = flag.String("name", "anonymous", "typeyour name")
fmt.Println(*data1)
// cara kedua 
var data2 string
flag.StringVar(&data2,"gender", "male", "type your gender")
fmt.Println(data2)
/*
Tinggal tambahkan akhiran var pada pemanggilan nama fungsi flag yang digunanan
Kegunaan dari parameter terakhir method-method flag adalah untuk memunculkan hints atau petunjuk 
arguments apa saja yang bisa dipakai, ketika 
argument -help ditambahakan saat eksekusi program.
*/