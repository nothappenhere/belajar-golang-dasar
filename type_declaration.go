package main

import "fmt"

func main(){
	type noKTP string

	var ktpAkbar noKTP = "11111111"

	var contoh string = "22222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpAkbar)
	fmt.Println(contohKTP)
}