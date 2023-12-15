package main

import "fmt"

func main() {
	example_func("Test")
	print(example_func_bool(0))

}

func example_func(example_string string){
	fmt.Println("Example string print: ", example_string)
}

func example_func_bool(example_int int) (bool){
	if  0 == example_int {
		return true
	}
	return false
}
