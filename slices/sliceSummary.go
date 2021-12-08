package main

import "fmt"

// peki ya en başta array boyutunu ne kadarlık bir değişkene sahip array tutacağımız  hakkında bilgimiz yok ise
// slice türünde diziye dinamik olarak sonradan elemanlar ekleyebiliyoruz .
// go da iki şekilde slice oluşturulmaya izin veriliyor bunlardan ilki make kullanımı iledir.


var arabalar = make([]int,5,10)

var motorsikletler = make([]int,7)

// 2. yol ise


func main() {



c:=cap(arabalar)
fmt.Println("arabalar slicenın kapasitesi ",c)
m:=cap(motorsikletler)
fmt.Println("motorsikletler kapasitesi",m)

	cars :=[]string{"mercedes","bmw"}
	fmt.Println(cars)

	// peki bir slicea nasıl eleman ekleriz ?
	// append fonksiyonu ile bir slicea eleman eklenir eğer yeni eleman deklare edilen kapasitenin üstünde ise önce kapasite artırılıp sonrasında eleman eklenecektir .

	arabalar = append(arabalar,12,1,8,95,23)
	arabalar = append(arabalar,45)
	fmt.Println("arabalar slice kapasite ",cap(arabalar))
	fmt.Println(arabalar)

}
