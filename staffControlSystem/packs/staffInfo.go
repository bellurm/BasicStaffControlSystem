package packs

import "fmt"

type staff struct {
	name     string
	age      int
	salary   float64
	pozition string
	id       int
}

var staffID int
var quit bool = false
var raise float64

// üst kademe yöneticiler
func Top_managers() []staff {
	fmt.Println("[*] Üst Kademe Yönetici Grubu Seçildi.")

	staff1 := staff{
		name:     "Aleda",
		age:      49,
		salary:   40000,
		pozition: "CEO",
		id:       1,
	}

	staff2 := staff{
		name:     "Cyber",
		age:      47,
		salary:   38000,
		pozition: "CFO",
		id:       2,
	}

	staff3 := staff{
		name:     "Worm",
		age:      46,
		salary:   38000,
		pozition: "COO",
		id:       3,
	}
	fmt.Println(staff1, staff2, staff3)
	for !quit {
		fmt.Print("Zam yapacaksanız çalışan ID, yapmayacaksanız '-1' giriniz: ")
		fmt.Scanln(&staffID)
		if staffID == -1 {
			break
		}

		fmt.Print("Zam oranını yazınız: ")
		fmt.Scanln(&raise)

		if staffID == staff1.id {
			fmt.Printf("%v Seçildi.\n", staff1)
			raise_rate := (float64(staff1.salary) * raise)
			staff1.salary = float64(staff1.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff1.salary)

		} else if staffID == staff2.id {
			fmt.Printf("%v Seçildi.\n", staff2)
			raise_rate := (float64(staff2.salary) * raise)
			staff2.salary = float64(staff2.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff2.salary)

		} else if staffID == staff3.id {
			fmt.Printf("%v Seçildi.\n", staff3)
			raise_rate := (float64(staff3.salary) * raise)
			staff3.salary = float64(staff3.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff3.salary)

		} else {
			fmt.Println("[*] Lütfen geçerli bir ID giriniz.")
		}
	}
	return []staff{staff1, staff2, staff3}
}

// orta kademe yöneticiler
func Mid_managers() []staff {
	fmt.Println("[*] Orta Kademe Yönetici Grubu Seçildi.")

	staff1 := staff{
		name:     "X",
		age:      30,
		salary:   22000,
		pozition: "Grup şefi",
		id:       4,
	}

	staff2 := staff{
		name:     "Y",
		age:      28,
		salary:   20000,
		pozition: "Grup şef yardımcısı",
		id:       5,
	}

	staff3 := staff{
		name:     "Z",
		age:      27,
		salary:   18500,
		pozition: "Çalışan",
		id:       6,
	}
	fmt.Println(staff1, staff2, staff3)
	for !quit {
		fmt.Print("Zam yapacaksanız çalışan ID, yapmayacaksanız '-1' giriniz: ")
		fmt.Scanln(&staffID)
		if staffID == -1 {
			break
		}

		fmt.Print("Zam oranını yazınız: ")
		fmt.Scanln(&raise)

		if staffID == staff1.id {
			fmt.Printf("%v Seçildi.\n", staff1)
			raise_rate := (float64(staff1.salary) * raise)
			staff1.salary = float64(staff1.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff1.salary)

		} else if staffID == staff2.id {
			fmt.Printf("%v Seçildi.\n", staff2)
			raise_rate := (float64(staff2.salary) * raise)
			staff2.salary = float64(staff2.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff2.salary)

		} else if staffID == staff3.id {
			fmt.Printf("%v Seçildi.\n", staff3)
			raise_rate := (float64(staff3.salary) * raise)
			staff3.salary = float64(staff3.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff3.salary)

		} else {
			fmt.Println("[*] Lütfen geçerli bir ID giriniz.")
		}
	}
	return []staff{staff1, staff2, staff3}
}

// alt kademe yöneticiler
func Low_managers() []staff {
	fmt.Println("[*] Alt Kademe Yönetici Grubu Seçildi.")

	staff1 := staff{
		name:     "A",
		age:      26,
		salary:   17000,
		pozition: "Grup şefi",
		id:       7,
	}

	staff2 := staff{
		name:     "B",
		age:      25,
		salary:   16000,
		pozition: "Grup şef yardımcısı",
		id:       8,
	}

	staff3 := staff{
		name:     "C",
		age:      24,
		salary:   14000,
		pozition: "Çalışan",
		id:       9,
	}
	fmt.Println(staff1, staff2, staff3)
	for !quit {
		fmt.Print("Zam yapacaksanız çalışan ID, yapmayacaksanız '-1' giriniz: ")
		fmt.Scanln(&staffID)
		if staffID == -1 {
			break
		}

		fmt.Print("Zam oranını yazınız: ")
		fmt.Scanln(&raise)

		if staffID == staff1.id {
			fmt.Printf("%v Seçildi.\n", staff1)
			raise_rate := (float64(staff1.salary) * raise)
			staff1.salary = float64(staff1.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff1.salary)

		} else if staffID == staff2.id {
			fmt.Printf("%v Seçildi.\n", staff2)
			raise_rate := (float64(staff2.salary) * raise)
			staff2.salary = float64(staff2.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff2.salary)

		} else if staffID == staff3.id {
			fmt.Printf("%v Seçildi.\n", staff3)
			raise_rate := (float64(staff3.salary) * raise)
			staff3.salary = float64(staff3.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff3.salary)

		} else {
			fmt.Println("[*] Lütfen geçerli bir ID giriniz.")
		}
	}
	return []staff{staff1, staff2, staff3}
}

// operasyonel çalışanlar
func Operational_staffs() []staff {
	fmt.Println("[*] Operasyonel Takım Grubu Seçildi.")

	staff1 := staff{
		name:     "A",
		age:      24,
		salary:   10000,
		pozition: "Çalışan",
		id:       10,
	}

	staff2 := staff{
		name:     "B",
		age:      23,
		salary:   10000,
		pozition: "Çalışan",
		id:       11,
	}

	staff3 := staff{
		name:     "C",
		age:      22,
		salary:   10000,
		pozition: "Çalışan",
		id:       12,
	}
	fmt.Println(staff1, staff2, staff3)
	for !quit {
		fmt.Print("Zam yapacaksanız çalışan ID, yapmayacaksanız '-1' giriniz: ")
		fmt.Scanln(&staffID)
		if staffID == -1 {
			break
		}

		fmt.Print("Zam oranını yazınız: ")
		fmt.Scanln(&raise)

		if staffID == staff1.id {
			fmt.Printf("%v Seçildi.\n", staff1)
			raise_rate := (float64(staff1.salary) * raise)
			staff1.salary = float64(staff1.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff1.salary)

		} else if staffID == staff2.id {
			fmt.Printf("%v Seçildi.\n", staff2)
			raise_rate := (float64(staff2.salary) * raise)
			staff2.salary = float64(staff2.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff2.salary)

		} else if staffID == staff3.id {
			fmt.Printf("%v Seçildi.\n", staff3)
			raise_rate := (float64(staff3.salary) * raise)
			staff3.salary = float64(staff3.salary) + raise_rate
			fmt.Printf("Çalışanın yeni Maaşı: %v \n", staff3.salary)

		} else {
			fmt.Println("[*] Lütfen geçerli bir ID giriniz.")
		}
	}
	return []staff{staff1, staff2, staff3}
}
