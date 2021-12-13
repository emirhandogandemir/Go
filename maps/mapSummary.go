package main

import "fmt"

// Map yapısı aslında bir koleksiyon yapısıdır . Aynı zamada bir sözlük gibi de düşünebiliriz. Bu yapı sayesinde key value değerleri ile kolay bir tanım sağlayabilmekteyiz

var sozluk map[string]string


var x map[int]string

func main() {

	sozluk = make(map[string]string)
	x := make(map[int]string)

	sozluk["go"]="popüler bir programlama dili"
	sozluk["java"]="oop tabanlı bir programlama dili"
	x[0]="sıfır"
	x[1]="bir"
	fmt.Println(sozluk,x)
}
