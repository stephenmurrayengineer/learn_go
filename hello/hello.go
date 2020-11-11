package main

// contains our Println function:
import "fmt"

// constanst are defined like this:
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "
const spanish = "Spanish"
const french = "French"
const italian = "Italian"

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}
	switch language {

	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	case italian:
		return italianHelloPrefix + name
	}
	return englishHelloPrefix + name

}

func main() {
	fmt.Println(Hello("World", ""))
}
