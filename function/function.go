package main

import (
	"fmt"
	"time"
)

func main() {
timer(time.After(5*time.Second))
}

//func timer(c <- chan time.Time){
//	for{
//		select {
//		case <- c:
//			return
//		default:
//			fmt.Println(time.Now())
//			time.Sleep(1*time.Second)
//		}
//	}
//}

func timer (c<- chan time.Time ,message ...string){
	defer fmt.Println("bir sonraki aşamaya geçiyor")
	defer fmt.Println("timer bitti")
	defer func () {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for  {
		select {
		case <- c:
			return
		default:
			fmt.Println(time.Now(),message)
			time.Sleep(1*time.Second)
			panic("panic atldı")
		}
	}

}



// 3 nokta ile ifade edilen string ler mesela istersen hiç gönderme istersen 100 tane gönder demek anlamına geliyor yani
// defer keywordu , function çağrısı bittikten sonra return ettikten sonra defer keywordune sahip olanlar sırası ile çağrılıyor
// uygulanmasını kapanmasını sağlayan panic function atılabilir , panicleri olabildiğince function içerisinde yukarılarda tanımlamak daha iyidir