package main

import (
	"errors"
	"fmt"
	"time"
)

type Kullanici struct { // Kullanici adinda bir struct olusturk.
	isim    string   `json:"isim"`    // Kullanicimiza string türünde bir isim adlı degisken verdik.
	soyIsim string   `json:"soyIsim"` // Kullanicimiza string türünde bir soyIsim adlı degisken verdik.
	begeni  []Begeni // Kullanicimiza Begeni türünde bir begeni adlı dizi verdik.
}

type Begeni struct { // Begeni adinda bir struct oluşturduk.
	id   string    // Begeni structimiza string türünde id adlı degisken verdik.
	date time.Time // Begeni structimiza time.Time türünde Date adlı degisken verdik.
}

func (k Kullanici) BegeniSayaci() int { // Begeni structimizin kullanabileceği içine kullanicidan bir degisken alan int türünde bir fonksiyon yazdık.
	return len(k.begeni) // kullanicimizin begeni adindaki dizisinin boyutunu döndürüyoruz.
}

func (k *Kullanici) Begeni() error { // Begeni Structimizin kullanabileceği Kullanicinin adresini tutan bir değişken alan error türünde bir fonksiyon tanımladik.
	if k.BegeniSayaci() >= 9 { // kullanicimizin begeni sayisi 9 dan büyükse kontrolü sağlanıyor.
		return errors.New("ulaşılan maksimum beğeni sayısı") // eger 9 dan büyükse uygulanan komut satiri.
	}

	if k.begeni == nil { // kullanicimizin begeni sayisi nil'mi(null) kontrolü yapiliyor.
		k.begeni = make([]Begeni, 0, 10) // eger nil'se(null) kullanicimizin begeni adli dizisi olusturuluyor.
	}

	b := Begeni{ // b degiskenine Begeni structı tanimlaniyor.
		id:   "id",
		date: time.Time{},
	}

	k.begeni = append(k.begeni, b) // kullanicimizin begenizi dizisine b ekleniyor.

	return nil
}

func main() {
	k1 := Kullanici{ // Kullanicimizi tanimladik.
		isim:    "Ktu",
		soyIsim: "Software",
	}

	for i := 0; i < 11; i++ { // 11 defa dönecek bir for döngüsü yazdik.
		if err := k1.Begeni(); err != nil { // err degiskenine kullanicimizin Begenisinin kontrolünün yapıldığı fonksiyonu tanimladik ve err nil değilse kontrolü yapiliyor.
			fmt.Println(err) // err yazdiriliyor.
			break            // ve iften çıkılıyor.
		}
		fmt.Println(k1.BegeniSayaci()) // kullanicinin begeni sayısı yazdiriliyor.
	}

}
