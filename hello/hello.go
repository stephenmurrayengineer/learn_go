package main

// contains our Println function:
import "fmt"

// constanst are defined like this:
const englishHelloPrefix = "Hello, "

// Capitalised function names:
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
