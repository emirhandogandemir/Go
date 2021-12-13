package main
import (
	"encoding/json"
	"fmt"
	// goda buit-in gelen kütüphane
	"time"
)
type Kullanıcı struct {
	Ad      string      `json:"adi"` // tag olarak tanımlanıyor
	Soyad   string      `json:"soyad"`
	Takipci []Kullanıcı `json:"takipci,omitempty"`
	//Begeni [] struct {
	//Tarih time.Time
	//}
	Begeni []Begen
}

type Begen struct {
	Tarih time.Time
}

func main() {
	k := Kullanıcı{Ad: "Emirhan", Soyad: "Doğandemir", Takipci: []Kullanıcı{
		{Ad: "Hasan", Soyad: "huseyin"}, {Ad: "Mehmet", Soyad: "Erbakan"},
	}}
	fmt.Println(k.Ad)
	arr, _ := json.Marshal(k) // bir structun json çıktısını almak istiyorsak bu paketi kullanmalıyız
	fmt.Println(string(arr))
	fmt.Println(k)
}
