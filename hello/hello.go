package main

// contains our Println function:
import "fmt"

// constanst are defined like this:
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == ""{
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
