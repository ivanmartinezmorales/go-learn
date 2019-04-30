package main

import "fmt"

func Greeting(name string) string {
	if name == "" {
		name = "Stranger"
	}
	return "Hello, " + name
}
func main() {
	fmt.Println(Greeting("Carlos"))
}
