package main

import "fmt"

const prefixHello string = "Hello, "

func Hello(name string) string {
	if name == "" {
		return prefixHello + "World"
	}

	return prefixHello + name
}

func main() {
	fmt.Println(Hello(""))
}
