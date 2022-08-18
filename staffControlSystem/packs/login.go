package packs

import (
	"fmt"
	"time"
)

func Login() {

	var quit bool = false
	username, password := "", "" //usr: CyberWorm | passwd: blog-cyberworm

	counter := 0 // 3 deneme sonrasında kilitlemek için

	for !quit {

		fmt.Print("Kullanıcı adı: ")
		fmt.Scanln(&username)

		fmt.Print("Parola: ")
		fmt.Scanln(&password)

		if username == "CyberWorm" && password == "blog-cyberworm" {
			fmt.Printf("Hoş geldiniz sayın %v. Çıkmak için '-1' komutunu giriniz.\n", username)
			break
		} else {
			counter++
			fmt.Println("[*] Kulllanıcı adı veya parola hatalı.")
			if counter == 3 {
				fmt.Println("[*] Çok sayıda deneme. Hesabınız kısa süreliğine engellendi.")
				time.Sleep(5 * time.Minute)
				quit = true
			}
		}
	}
}
