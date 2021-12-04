package main

import (

	"fmt"
)

func main() {

	slc_1 := [] int{1,2,3,4,5}
	slc_2 := []int{}

	for i ,value :=range slc_1{
		fmt.Println("index %d value: %d \n",i,value)
	}

	for i := 0; i<10 ;i++ {
		slc_2 =append(slc_2,1)
	}

}
