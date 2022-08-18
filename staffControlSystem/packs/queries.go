package packs

import (
	"fmt"
)

func Query() {

	var quit bool = false
	staff_levels := [5]int{-1, 1, 2, 3, 4}
	var choice int

	for !quit {

		fmt.Print("Hangi kademeden bilgi istiyorsunuz?\n1.Üst kademe\n2.Orta kademe\n3.Alt kademe\n4. Operasyonel kademe\n> ")
		fmt.Scanln(&choice)
		if choice == staff_levels[0] {
			fmt.Println("Sistemden çıkıldı.")
			quit = true
		} else if choice == staff_levels[1] {
			fmt.Println(Top_managers())
		} else if choice == staff_levels[2] {
			fmt.Println(Mid_managers())
		} else if choice == staff_levels[3] {
			fmt.Println(Low_managers())
		} else if choice == staff_levels[4] {
			fmt.Println(Operational_staffs())
		} else {
			fmt.Println("[*] Geçerli bir kademe giriniz.")
		}
	}

}
