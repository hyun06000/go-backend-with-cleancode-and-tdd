package main

import "fmt"

const prefixHello string = "Hello, "

func Hello(name string) string {
	if name == "" {
		return prefixHello + "World"
	}

	return prefixHello + name
}

func hellotest() { // main
	fmt.Println(Hello(""))
}
