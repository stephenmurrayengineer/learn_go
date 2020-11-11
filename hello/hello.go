package main

// contains our Println function:
import "fmt"

// Capitalised function names:
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
