package main

import (
	"fmt"
	"strings"
)

type String * string // reference type tanımı yapıldı burada

func main() {

	var rstr String
	var pstr string
	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "go türkiye"
	rstr = &pstr
	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "go turkiye egitim kampı"

	fmt.Println(*rstr)
	fmt.Println(pstr)
	pstr = "go turkiye egitim kampı 2021"
	//replaceStr(pstr)
	//fmt.Println(*rstr)
	//fmt.Println(pstr)
	replaceStr(rstr)
	fmt.Println(*rstr)
	fmt.Println(pstr)

}
//func replaceStr(s string){
//	s = strings.Replace(s,"go","go",1)
//}

func replaceStr(s String){
	*s= strings.Replace(*s,"go","GO",1)
}