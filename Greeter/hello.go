package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "Kon'nichiwa, "
const spanish = "Spanish"
const french = "French"
const japanese = "Japanese"

// Hello - Function to return greeting
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name + "!"
}

func getPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
