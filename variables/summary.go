package main

import (
	"fmt"
)

// değişken tanımlarımız için var keywordu ile fonksiyon içerisinde olmaksısızn değişken tanımı yapılabilirken := operatörü ile fonksiyon içerisinden çağrılarımızı yapabiliriz .



var dogruMu bool // veri tipi belirtilerek yalnızca tanımlama çeşitleri
var yas int
var adı string

var boy int = 2 // değer atarken veri tipini de belirterek kullanma
var soyad string ="Emirhan"

var x ="nasılsın abi" // veri tipi belirtmeden kullanımı artık x değişkeni string veri türünde değere sahiptir

var c,c2,c3 = "naber","nasılsın",12 // tek satırda birden fazla data type belirtmeden değişken tanımları

var k,l,m int =1,2,3 // data type belirterek tek satırda değişken tanımları burada tüm değişkenler aynı tipte olmak zorunda

//var w,q,z string = "ahmet","hasan",3

// q:="selam abi"


//constların farklılığı ise sabitlerin değerini tanımlama yaparken atamamız gerekliliği ve atandıktan sonra değiştirilemez olması
const variable1 ="hasan"
const variable =18

const (
	variable3 ="deger3"
	variable4="deger4"
)

const (
	variable5=iota
	variable6
	variable7
	variable8
)


func main() {

	// default zero values konusu

//	var fathername string = "ramazan" var ilke function içerisinde de tanım yapılabiliyor konuda herhangi bir olumsuzluk yok

	ktü := "go türkiye"
	swe:=12
fmt.Println(x, c3,m,ktü,swe)

}
