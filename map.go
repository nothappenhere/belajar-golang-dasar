package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "Rizky Akbar",
		"address" : "karawang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Rizky Akbar"
	book["wrong"] = "Ups"
	
	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}