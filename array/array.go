package main

import (
	"fmt"
)

var arr_1 [3]string
var arr_2 = [6]int{1, 2, 3, 4, 5, 6}

func main() {

	//arr_3 := make([]int, 8)

	//arr_3[1] = 18


	programmingLanguage := [6]string{"java", "go", "c++", "c", "python"}

	fmt.Println(programmingLanguage)
	//fmt.Println(arr_1, arr_2, arr_3)
	fmt.Println("arr_1 len \n", len(arr_1))
	//fmt.Println("arr_2 len \n", len(arr_3))
}
