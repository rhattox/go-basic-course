package main

import "fmt"

func main() {
	example_string := "String Example"
	fmt.Println(example_string)

	var example_int int
	example_int = 13
	fmt.Println(example_int)

	var example_float float32
	example_float = 1.234
	fmt.Println(example_float)

	example_bool := true
	fmt.Println(example_bool)

	example_string_array := [3]string{"banana","apple","grape"}

	for _, fruits := range example_string_array {
		fmt.Println(fruits)
	}

	example_int_array := [...]int{17,22}

	for _, number := range example_int_array{
		fmt.Println(number)
	}

}