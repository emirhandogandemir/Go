package main

//Bildiğiniz gibi, her değişken bir bellek konumudur ve her bellek konumunun,
//bellekteki bir adresi gösteren ve ve işareti (&) işleci kullanılarak erişilebilen bir adresi tanımlıdır.
//Pointer, başka bir değişkenin bellek adresini depolayan bir değişkendir.

import "fmt"

type String *string

func main() {
	var value_1 String // burada normal String den bir değişken tanımladık.
	var value_2 string // yukarıda tanımladığımız String türünden bir değişkenden bir değişken tanımladık

	fmt.Println(value_1) // value_1 degerimizi yazdırdık.
	fmt.Println(value_2) // value_2  degerimizi yazdırdık
	println("***************************************************************************")
	value_2 = "Ktu Software Engineering" // value_2 degiskenine bir deger atadık.
	value_1 = &value_2                   // value_1 degiskenine value_2 nin bellekte tutulan adresini atadık.

	fmt.Println(value_1) // value_1 degerimizi yazdırdık. yani value_2 nin bellekteki adresini göreceğiz.
	fmt.Println(value_2) // value_2  degerimizi yazdırdık.
	fmt.Println(&value_2)
	fmt.Println(*value_1) // * ile bellekteki adresin içindeki degeri gösterdik.
	println("***************************************************************************")

	value_2 = "Go Bootcamp" // value_ nin degerini değiştirelim.

	fmt.Println(value_1)  // value_1 i ekrana yazdıralım.
	fmt.Println(*value_1) // value_1 in tuttuğu bellek adresindeki degeri ekrana yazdıralım
	fmt.Println(value_2)  // value_2 yi ekrana yazdıralım.
	fmt.Println(&value_2)
	println("***************************************************************************")
	*value_1 = "Ktu Software Engineering Go Bootcamp"

	fmt.Println(*value_1)
	fmt.Println(value_2)
	println("***************************************************************************")

}
