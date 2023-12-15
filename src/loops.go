package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 0
	for counter < 5 {
		fmt.Println("Counter:", counter)
		counter++
	}

	colors := [3]string{"Red", "Green", "Blue"}

	for index, value := range colors {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	person := map[string]string{"Name": "John", "Age": "30", "Phone" :"23456789"}

	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	message := "Hello, Zina!"

	for _, char := range message {
		fmt.Printf("%c ", char)
	}
}
