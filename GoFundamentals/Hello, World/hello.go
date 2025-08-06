package main

import "fmt"

const SPANISH = "Spanish"
const FRENCH = "French"

const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func HelloWorld() string {
	return englishHelloPrefix + "World!"
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {

	case FRENCH:
		return fmt.Sprintf("%s%s!", frenchHelloPrefix, name)
	case SPANISH:
		return fmt.Sprintf("%s%s", spanishHelloPrefix, name)
	default:
		return fmt.Sprintf("%s%s!", englishHelloPrefix, name)
	}
}

func main() {
	fmt.Println(HelloWorld())
}
