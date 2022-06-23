package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	var index, value = min(x)
	fmt.Println("Список значений : ", x)
	fmt.Printf("Минимальное значение в x[%d] = %d\n", index, value)
}

func min(array []int) (min_index int, min_value int) {
	min_index = 0
	min_value = array[min_index]
	for i, v := range array {
		if v < min_value {
			min_value = v
			min_index = i
		}
	}
	return
}
