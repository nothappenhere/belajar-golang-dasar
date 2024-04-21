package main

import "fmt"

func main(){
	name := "Rizky"

	switch name {
	case "Akbar":
		fmt.Println("Hallo", name)
	case "Rizky":
		fmt.Println("Hallo?")
	default:
		fmt.Println("Kamu Siapa?")
	}

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("Namanya pas")
	case false:
		fmt.Println("Namanya Kurang panjang nih")
	}

	name = "Rizky Akbar"
	length := len(name)
	switch  {
	case length > 10:
		fmt.Println("Namanya kepanjangan")
	case length >= 5:
		fmt.Println("Namanya pas")
	default:
		fmt.Println("Situ punya nama?")
	}
}