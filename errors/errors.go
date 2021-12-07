package main

import (
	"errors"
	"fmt"
	"log"
)

type Kullanici struct{
	Ad string `json:"adi"`
	Soyad string `json:"soyad"`
	Takipci []Kullanici
}

func (k Kullanici) TakipciSayisi() int {
	return len(k.Takipci)
}

func (k *Kullanici) TakipciEkle(t Kullanici) error{
	if k.TakipciSayisi()==10 {
		return errors.New("maximum takipci sayısına ulaşılmıştır")
	}
	if k.Takipci==nil {
		k.Takipci= make([]Kullanici,0,10)
	}
	k.Takipci = append(k.Takipci,t)
	return nil
}

func main() {

	k:=Kullanici{
		Ad: "GO",
		Soyad: "Turkiye",
		Takipci: []Kullanici{
			{
				Ad: "hasan",
				Soyad: "1",
			},
			{
			Ad: "Mehmet",
			Soyad: "2",
			},
		},
	}
	t:=Kullanici{
		Ad: "Takipci",
		Soyad: "-",
	}
	fmt.Println(k.TakipciSayisi())
	if err :=k.TakipciEkle(t); err!=nil{
		log.Fatal(err)
	}
	fmt.Println(k.TakipciSayisi())

}
