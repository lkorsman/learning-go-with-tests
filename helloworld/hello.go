package main

import "fmt"


const (
	finnish = "Finnish"
	french = "French"
	spanish = "Spanish"
	englishHelloPrefix = "Hello, "
	finnishHelloPrefix = "Hei, "
	frenchHelloPrefix = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)



func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case finnish:
		prefix = finnishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}