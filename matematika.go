package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println("Penjumlahan a + b =",c)

	var aa = 10
	var bb = 5
	var cc = aa - bb
	fmt.Println("Pengurangan aa + bb =",cc)

	var aaa = 5
	var bbb = 10
	var ccc = aaa * bbb
	fmt.Println("Perkalian aaa + bbb =",ccc)

	var aaaa = 10
	var bbbb = 5
	var cccc = aaaa / bbbb
	fmt.Println("Pembagian aaaa + bbbb =",cccc)

	var i = 10
	i += 10
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}