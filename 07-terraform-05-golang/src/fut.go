package main

import "fmt"

func main() {
	fmt.Print("Enter a meters: ")
	var m float64
	fmt.Scanf("%f", &m)

	//output := input * 3.28

	fmt.Println("Footage: ", convert(m))
}

func convert(m float64) float64 {
	return m * 3.28
}
