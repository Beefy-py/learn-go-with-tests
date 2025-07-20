package main

import "fmt"


func HelloWorld() string {
	return "Hello, World!";
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, e%s!", name)
}

func main() {
	fmt.Println(HelloWorld())
}