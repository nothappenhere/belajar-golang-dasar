package main

import "fmt"

func main(){
	name := "Riz"

	if (name) == "Akbar"{
		fmt.Println("Hallo", name)
	}else if (name == "Rizky Akbar")  {
		fmt.Println("America yaa")
	}else {
		fmt.Println("Siapa Kamu?")
	}

	if lenght := len(name); lenght > 5{
		fmt.Println("Nama Nya Keren")
	}else {
		fmt.Println("det sit?")		
	}
}