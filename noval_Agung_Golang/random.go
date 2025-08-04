package novalagunggolang

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func random() {
	//Pada chapter ini kita akan belajar pemanfaatan package math/rand untuk pembuatan data acak atau random
	ranomizer := rand.New(rand.NewSource(10))
	fmt.Println("random ke 1:", ranomizer.Int())
	fmt.Println("random ke 2:", ranomizer.Int())
	fmt.Println("random ke 3:", ranomizer.Int())

	/*S*/
}

func exampleMain() {
	//Casting string -> byte
	/*string sebenarnya slice/array byte.
	Di go sebuah karakter biasa (bukan uniqode)
	direpresentasikan oleh sebuah elemen slice byte.Tiap elemen data bersis data integer dengan basis desimal yang
	merupakan kode ASCII dari dalam string
	cara mendapatkan slice byte dari sebuah data string adalah dengan mengcastingnya ke type
	*/
	var text1 = "hallo"
	var b = []byte(text1)
	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	var byte = []byte{104, 97, 108, 111}
	var s = string(byte)
	fmt.Printf("%s \n ", s)

	var c int64 = int64('h')
	fmt.Println(c)
	var d string = string(104)
	fmt.Println(d)

	//Penerapan Regexp
	var text = "banana burger soup"
	var regex1, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex1.FindAllString(text, 2)
	fmt.Println("%#v \n", res1)
	var res2 = regex1.FindAllString(text, 1)
	fmt.Println("%#v \n", res2)

	//Method matchstring
	var text2 = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)
	var isMath = regex.MatchString(text2)
	fmt.Println(isMath)

	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

}

func MethodFindString() {
	/*Digunakan untuk mencari index string kembalian hasil dari operasi regexp*/
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)

	var str = text[0:6]
	fmt.Println(str)
	/*Method ini sama dengan FindString() hanya saja yang dikembalikan indeksnya */
}

func FindAllString() {

}

/*Encode - Decode Base64
Go menyediakan package encoding/base64
berisikan fungsi-fungsi untuk mkebutuhan encode dan decode data kebentuk base64 dan sesbaliknya.
data yang akan di encode bertipe []byte, maka perlu dilakukan casting untuk
data-data yang tipenya belum []byte.
Proses encoding dan decoding bisa dilakukan dan via bebebrapa caara yang pada chapter ini akan pelajari


*/

func EncodeAndDecode() {
	var name = "Sukma Rizki"
	var encodeString = base64.StdEncoding.EncodeToString([]byte(name))
	fmt.Println("encoded:", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodedString = string(decodeByte)
	fmt.Println("decoded:", decodedString)

	var data = "John F khanedy"
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodeString1 = string(encoded)
	fmt.Println(encodeString1)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())

	}
	var decodedString1 = string(decoded)
	fmt.Println(decodedString1)

}

func EncodeDecodeUrl() {
	var dataUrl = "https://Kalipare.com"

	var encodestring = base64.URLEncoding.EncodeToString([]byte(dataUrl))
	fmt.Println(encodestring)

	var decodebyte, _ = base64.URLEncoding.DecodeString(encodestring)
	var decodeString = string(decodebyte)
	fmt.Println(decodeString)

}

func HashSha1() {
	var text = "this is scaret"
	var sha = sha1.New()
	sha.Write([]byte(text))

	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
	/*Metode Salting pada Hash SHA1
	  Salt dalam konteks kriptografi adalah data acak yang digabungkan pada data aslu
	  sebelum proses hash dilakukan
	  Hash merupakan enkripssi satu arah dengan lebar data yang sudah pasti, sangat mungkin sekali
	  kalau hasil hash unutuk beberapa data adalah sama. Disinilah kegunaan salt

	*/

}

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encryted = sha.Sum(nil)
	return fmt.Sprintf("%x", encryted), salt
}

func MainHash() {
	var text = "this is secret"
	fmt.Printf("original : %s\n \n ", text)
	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1: %s\n \n", hashed1)
	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2: %s \n \n", hashed2)
	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3: %s \n \n", hashed3)
	_, _, _ = salt1, salt2, salt3
}
