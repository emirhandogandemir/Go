package main

//Döngü ifadesi, bir ifadeyi veya ifade grubunu birden çok kez yürütmemize izin verir
//ve programlama dillerinin çoğunda bir döngü deyiminin genel biçimini takip eder

import (
	"fmt"
)

func main() {
	// 2 tane slice tanımladım. bir tanesi 3 elamanlı diğerinin elemanı bulunmamakta.
	slc_1 := []int{1, 2, 3}
	slc_2 := []int{}

	// burada for döngüsünü kullanara bir foreach döngüsü oluşturdum. Go da Foreach döngüsünün oluşturulabilir olabilmesi range komutu na dayanmakta.
	// range komutu , yanına girilen array yada slice dizilerinin elemanlarını gezmeye yarıyor.
	for i, value := range slc_1 { // i foreach döngüsünün indexi value ise indexin degeri biz burada value degerine slc_1 adlı dizimizin elemanlarını atıyoruz.
		fmt.Printf("index : %d value :%d \n", i, value) // %d integer verilerin yazdırılmasında kullanılır. yani indexten sonra %d koyduğumda int bir deger vereceğimi söylüyorum.
	}

	// buradaki for döngümüzde ise clasic bir for döngüsü görmekteyiz.
	// şuana kadar gördüğümüz syntax yapılarında ; görmedik fakat for döngüsünün yapısında kullanıyoruz.
	// i ye 0 degerini atadık ve i 10 dan küçük olduğu sürece işlemimizi gerçekleştiriyoruz. ve her dönmesinde i yi 1 artırıyoruz.
	// işlemimizde ise slc_2 dizimize i degerini ekliyoruz.
	for i := 0; i < 10; i++ {
		slc_2 = append(slc_2, i)
	}

	fmt.Println("slc_1 =>", slc_1) // slc_1 dizimizin çıktısını alıyoruz.
	fmt.Printf("slc_1 => len : %d , cap : %d \n", len(slc_1), cap(slc_1))
	fmt.Println("slc_2 =>", slc_2) // slc_2 dizimizin çıktısını alıyoruz.
	fmt.Printf("slc_2 => len : %d , cap : %d \n", len(slc_2), cap(slc_2))

	/*for _, value_2 := range slc_2 {
		for _, value_1 := range slc_1 {
			value_1 += value_2
		}
	}*/

	for _, value_2 := range slc_2 { // value_2 ye slc_2 nin elemanlarını atıyoruz
		for i := range slc_1 { // i ye slc_1 in elemanlarını atıyoruz.
			slc_1[i] += value_2 // burada slc_1 in i. indexindeki elemanı ile value_2 yi topluyoruz. Böylelikle indexteki eleman slc_2 nin elemanlarının toplamı ile toplanmış oluyor.
		}

	}

	fmt.Println("slc_1 + slc_2 =", slc_1)

	// Go da tek döngü yapısı bulunmakta buda for döngüsüdür.

	/*c := time.After(5 * time.Second) //5 saniyelik bir timer.

	// for yazıp süslü parantez açıp kapatırsak sonsuz döngü oluşturmuş oluruz.
	for {
		boolean := false
		select {
		case <-c:
			boolean = true //5 saniye sonra boolean true
		default:
			fmt.Println(time.Now()) //saniyede bir ekrana tarihi yazdırıyor.
			time.Sleep(1 * time.Second)
		}

		if boolean {
			break
		}
	}*/

}
