package main

import "fmt"

func sumup(val int) int {
	if val == 1 {
		return val
	}
	return val + sumup(val-1)
}

func main() {
	ret := sumup(10)
	fmt.Println("ret = ", ret)
}
