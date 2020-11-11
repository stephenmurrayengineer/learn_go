package main

// contains our Println function:
import "fmt"

// constanst are defined like this:
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}
	if language == spanish{
		return spanishHelloPrefix + name
	}
	if language == french{
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
