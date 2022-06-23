package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	index := 0
	fmt.Println("Список значений : ", x)
	for n, value := range x {
		if n == 0 {
			index = value
		} else {
			if value < index {
				index = value
			}
		}
	}
	fmt.Println("Минимальное число : ", index)
}
