package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const spanish = "Spanish"
const frenchPrefix = "Bonjour, "
const french = "French"

func Greeting(name string, language string) string {
	if name == "" {
		name = "Stranger"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Greeting("Carlos", "Spanish"))
}
