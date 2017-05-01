package main

import "fmt"

const pi = 3.1415

func main() {
	const World = "World!"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "day")

	const Truth bool = true
	fmt.Println("Go rules?", Truth)
}
