package main

// contains our Println function:
import "fmt"

// constanst are defined like this:
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}
	if language == "Spanish"{
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
