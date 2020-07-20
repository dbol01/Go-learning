package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello - Function to return string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("World"))
}
