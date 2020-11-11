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

// in go, public functions begin with Capital letters, private functions are lowercase

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name

}

// second set of brackets defines return string 
// this is assigned the 'zero' value: ""
func greetingPrefix(language string) (prefix string){
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
