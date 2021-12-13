package main

import "fmt"

// go programlama dilinde classlar yerine structlar vardır . Struct yapıları sayesinde bir nesne oluşturabilir ve bu nesneye ait özellikler oluşturabiliriz .

// type ile yeni bir tür oluşturup türünün de struct olduğunu söylüyoruz .

type Araba struct {
	marka string
	renk string
	maximumHız int

}

type Person struct {
	Id int
	name string
	age int
	surname string
}
 type Employee struct {
	 Person
	 ManagerId int
 }

func main() {
araba1 := Araba{"Bmw","siyah",250}
araba1.renk="pembe"
araba1.maximumHız=100
araba1.marka="audi"
fmt.Println(araba1)

employee1:= Employee{
	Person{Id: 1,name: "ahmet",age: 34,surname: "hasan"},
	12,
}
fmt.Println(employee1)
employee1.nameTell()
}

func (person *Person)nameTell(){
	fmt.Println(person.name)
}