package main

import "fmt"

// Arrayler sabit boyuta ve aynı veri tipine sahiptirler

var arr1 [5]int // şeklinde 5 boyutlu int veri tipinden oluşan bir array ifadesi yapılabilir .

// bir dizi bildiririrken mutlaka boyutu belirtilmeli . dizi tanımlandıktan sonra boyutu değiştirilemez

var arr2 = [6]string{"ahmet", "hasan", "necmi", "", "", ""} // bu da array oluştururken degerlerin de atandığı bir sözdizimidir

func main() {
	arr1[4] = 12

	arr3 := make([]int, 4) // eger fonksiyon içerisinde bir tanım yapacaksanız make keywordu ile kullanımı şekildeki gibidir

	//programmingLanguage := [6]string {"java","go","c++","c","python"} make kullanmadan kullanımı ise şu şekildedir

	arr3[2] = 12

	fmt.Println(arr1, arr3)
	fmt.Println("array2 boyutu", len(arr2)) // built-in gelen len functionu içerisinde arraylerin uzunluğunu bu şekilde bulabiliyoruz
}
