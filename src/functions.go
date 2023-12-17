package main

import "fmt"

func main() {
	example_func("Test")
	print(example_func_bool(0))
	example_multi_return("oi", true)
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

func example_multi_return(example_string string, example_bool bool) (string, bool){
	fmt.Println(example_string, example_bool)
	return example_string,example_bool
}