package main

import "fmt"

func main() {
	bar()
	foo()
}

func bar() {
	fmt.Println("bar")
}

func foo() {
	fmt.Println("foo")
}
