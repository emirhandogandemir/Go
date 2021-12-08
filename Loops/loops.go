package main

import (
	"fmt"
	"time"
)

func main() {

	slc_1 := [] int{1,2,3,4,5}
	slc_2 := []int{}

	for i ,value :=range slc_1{
		fmt.Println("index value:  \n",i,value)
	}

	//for i := 0; i<10; i++ {
	//	slc_2 =append(slc_2,1)
//	}

for i := range [10]int{}{
	slc_2 =append(slc_2,i)
}

for _, value_2 := range slc_2 {
	for _, value_1 := range slc_1 {
		value_1 +=value_2
	}
}

for _, value2 := range slc_2 {
	for i := range slc_1 {
		slc_1[i] += value2
	}
}


c := time.After(5*time.Second)

for{
	b := false
	select {
	case <-c:
		b= true
	default:
		fmt.Println(time.Now())
		time.Sleep(1*time.Second)
	}
	if b{
		break
	}
}
fmt.Println("test")



	fmt.Println(slc_1)
	fmt.Println(slc_2)
	fmt.Println("slc1 len and cap",len(slc_1),cap(slc_1))
	fmt.Println("slc2 len and cap",len(slc_2),cap(slc_2))
}
