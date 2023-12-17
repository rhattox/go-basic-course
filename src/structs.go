package main

import "fmt"

type Person struct {
	Name string
	Age int
	Height string
	Weight int
}

func main() {

	examplePerson := Person{
		Name: "Bruno",
		Age: 27,
		Height: "1.82",
		Weight: 70,
	}

	fmt.Println(examplePerson)
}
