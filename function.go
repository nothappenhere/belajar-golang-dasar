package main

import "fmt"

func helloWorld(){
	fmt.Println("Hello World!")
}

func sayMyName(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func myName(name string) string{
	hello := "Hello " + name
	return hello
}

func fullName() (string, string){
	return "Rizky", "Akbar"
}

func completeName() (firstName, middleName, lastName string){
	firstName = "Muhammad"
	middleName = "Rizky"
	lastName = "Akbar"

	return firstName, middleName, lastName
}

func main(){
	helloWorld()
	sayMyName("Rizky", "Akbar")

	result := myName("Akbar")
	fmt.Println(result)
	fmt.Println(myName("Rizky"))

	// firstName, lastName := fullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := fullName()
	fmt.Println(firstName)

	a, b, c := completeName()
	fmt.Println(a, b, c)
}