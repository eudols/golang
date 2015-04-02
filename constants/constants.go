package main

import "fmt"

const Pi = 3.1415

func main() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth bool = true
	fmt.Println("Go rules?", Truth)
}
