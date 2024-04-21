package main

import "fmt"

func main(){
	// i := 10
	// for i >= 0 {
	// 	fmt.Println("Nilai perulangan saat ini:", i)
	// 	i--
	// }
	// fmt.Println("Selesai")

	// for counter := 0; counter <= 10; counter++{
	// 	fmt.Println("Nilai counter:", counter)
	// }

	names := []string{"Muhammad", "Rizky", "Akbar"}
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	for index, name := range names{
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names{
		fmt.Println(name)
	}
}