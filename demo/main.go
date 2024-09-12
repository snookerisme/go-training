package main

import "fmt"

func main() {
	fmt.Println("Hello, Eiei!")

	resultSum := calculate(sum, 1, 2)
	fmt.Println(resultSum)

	resultMinus := calculate(minus, 4, 2)
	fmt.Println(resultMinus)

}

type myFunc func(int, int) int

func calculate(fn myFunc, a, b int) int {
	return fn(a, b)
}

func sum(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}
