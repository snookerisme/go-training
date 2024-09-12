package main

import "fmt"

type Flight struct {
	number      int
	airlineCode string
	departure   string
}

func main() {

	f := Flight{}
	fmt.Println("%#v\n", f)

	f1 := Flight{123, "AA", "123"}
	fmt.Println("%#v\n", f1)
	fmt.Println("airlineCode : ", f1.airlineCode)
	fmt.Println("number :", f1.number)

	f2 := Flight{airlineCode: "AA", number: 123}
	fmt.Println("%#v\n", f2)

}
