package main

import (
	"fmt"
	"strings"
)

func main() {

	str :="ktü go bootcampine hoşgeldiniz"
	str_1 :=str[:7]
	str_2 :=str[len(str)-4:]
	str_3:= fmt.Sprint(str_1,str_2)
	if strings.EqualFold(str_1,"LOrEM"){
		fmt.Println("str_1 equal to LOrEM")
	}




	fmt.Println(str)
	fmt.Println(str_1)
	fmt.Println(str_2)
	fmt.Println(str_3)
	fmt.Println(strings.ToUpper(str))
}
