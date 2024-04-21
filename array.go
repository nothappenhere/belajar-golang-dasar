package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Rizky"
	names[2] = "Akbar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{80, 90, 95}
	fmt.Println(values)

	var nilai = [...]int{
		50,
		40,
		30,
	}
	fmt.Println(len(nilai))

}