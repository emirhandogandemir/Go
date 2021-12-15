package main

//Fonksiyon, birlikte bir görevi gerçekleştiren bir dizi deyimdir.
//Fonksiyon bildirimi, derleyiciye bir Fonksiyon adı, dönüş türü ve parametreler hakkında bilgi verir.
//Fonksiyon tanımı, Fonksiyonun gerçek gövdesini sağlar

import "fmt"

/*func function_name( [parameter list] ) [return_types]
{
body of the function
}*/

func Topla(x, y int) int { // int türünde Topla adında bir fonksiyon oluşturalım ve bu fonksiyonumuza 2 deger girdisi olsun.
	return x + y // burada bu fonsiyon bize gönderdiğimiz 2 int degeri toplayıp gönderiyor.
}

func Buyuk(x int, y int) int { // int türünde Buyuk adında bir fonksiyon oluşturalım ve bu fonksiyonumuza 2 deger girdisi olsun.
	var value int // Büyük olan degeri tutmak için oluşturulmuş local değişken.
	if x > y {    // x büyük mü y den kontrolü sağlayan if satırı
		value = x // x , y den büyükse yapılacak olan işlem.
	} else { // eger x , y den büyük değilse kontrolü sağlayan else satırı
		value = y // x , y den büyük değilse yapılacak olan işlem.
	}

	return value // kontrol sonucunda atanılan degeri döndürüyoruz.
}

func Yazdir(deger, deger_1 string) (string, string) { // string türünde Yazdir adlı bir fonksiyon.

	return deger_1, deger // Bu fonksiyon 2 deger birden döndürdüğü için fonksiyonu tanımlarken 2 kere string ifadesi yazdık.
}

func Yazdir_1(deger string) string { // string türünde Yazdir_1 adlı fonksiyon.

	return deger // burada tek deger döndürdüğümüz için yukarıda fonksiyon tanımında tek string ifadesi yazdık.
}

func Helloworld() { // Boş bir fonksiyon tanımlandı.
	fmt.Println("Helllo , World!") // yapacak olduğu komut satırı.
}

func main() {
	var a int = 100
	var b int = 200

	var x string = "Ktu"
	var y string = "Software"

	var value int
	value = Buyuk(a, b)

	fmt.Printf("Buyuk Function : %d \n", value)
	fmt.Printf("Buyuk Function : %d \n", Buyuk(a, b))
	value = Topla(a, b)
	fmt.Printf("Topla Function : %d \n", value)
	fmt.Printf("Topla Function : %d \n", Topla(a, b))

	k, w := Yazdir("Go", "Bootcamp")
	fmt.Println("Yazdir Function : ", k, w)

	c, d := Yazdir(x, y) // burada 2 string ifadesi kullandığımız fonksiyonu hatırlayalım.
	// c ve d bizim o string ifadelerimizin karşılığı oluyor.
	fmt.Println("Yazdir Function : ", c, d)

	s := Yazdir_1(y)
	fmt.Printf("Yazdir_1 Function :%s \n", Yazdir_1(x))
	fmt.Printf("Yazdir_1 Function :%s \n", s)

	Helloworld()
}
