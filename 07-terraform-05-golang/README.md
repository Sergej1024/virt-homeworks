# Домашнее задание к занятию "7.5. Основы golang"
## Задача 1. Установите golang.
1. Воспользуйтесь инструкций с официального сайта: [https://golang.org/](https://golang.org/).
2. Так же для тестирования кода можно использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

```shell
[sergej@fedora Загрузки]$ sudo tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
[sergej@fedora Загрузки]$ export PATH=$PATH:/usr/local/go/bin
[sergej@fedora Загрузки]$ go version
go version go1.18.3 linux/amd64
[sergej@fedora Загрузки]$
```

## Задача 2. Знакомство с gotour.
У Golang есть обучающая интерактивная консоль [https://tour.golang.org/](https://tour.golang.org/).
Рекомендуется изучить максимальное количество примеров. В консоли уже написан необходимый код,
осталось только с ним ознакомиться и поэкспериментировать как написано в инструкции в левой части экрана.  

> Ознакомился

## Задача 3. Написание кода.
Цель этого задания закрепить знания о базовом синтаксисе языка. Можно использовать редактор кода
на своем компьютере, либо использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

1. Напишите программу для перевода метров в футы (1 фут = 0.3048 метр). Можно запросить исходные данные
у пользователя, а можно статически задать в коде.
    Для взаимодействия с пользователем можно использовать функцию `Scanf`:
    ```
    package main

    import "fmt"

    func main() {
        fmt.Print("Enter a number: ")
        var input float64
        fmt.Scanf("%f", &input)

        output := input * 2

        fmt.Println(output)    
    }
    ```

> [формула расчета](https://www.metric-conversions.org/ru/length/meters-to-feet.htm)

  ```go
    package main

    import "fmt"

    func main() {
    	fmt.Print("Enter a meters: ")
    	var input float64
    	fmt.Scanf("%f", &input)

    	output := input * 3.28

    	fmt.Println("Footage: ", output)
    }
  ```
>  При выполнении 4 задания возникла необхоимость в использовании именнованой функции.

  ```go
    package main

    import "fmt"

    func main() {
    	fmt.Print("Enter a meters: ")
    	var m float64
    	fmt.Scanf("%f", &m)

    	//output := input * 0.3048

    	fmt.Println("Footage: ", convert(m))
    }

    func convert(m float64) float64 {
    	return m * 3.28
    }
  ```

    <p align="center">
      <img width="1200" src="./img/fut.png">
    </p>

1. Напишите программу, которая найдет наименьший элемент в любом заданном списке, например:
    ```
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    ```

```go
package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	n := 0
	fmt.Println("Список значений : ", x)
	for i, v := range x {
		if i == 0 {
			n = v
		} else {
			if v < n {
				n = v
			}
		}
	}
	fmt.Println("Минимальное число : ", n)
}
```

> При выполнении 4 задания возникла необхоимость в использовании именнованой функции.

```go
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
```

<p align="center">
  <img width="1200" src="./img/test.png">
</p>


1. Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. То есть `(3, 6, 9, …)`.

```go
package main

import "fmt"

func main() {
	var n int
	var m int
	n = 1
	m = 100

	for i := n; i <= m; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
		}
	}
}
```
<p align="center">
  <img width="1200" src="./img/test2.png">
</p>

## Задача 4. Протестировать код (не обязательно).

Создайте тесты для функций из предыдущего задания.

```go
package main

import "testing"

func TestConvert(t *testing.T) {
	expected := 3.28
	received := convert(1)
	if received != expected {
		t.Errorf("Error, got: %f, want: %f", received, expected)
	}
}
func TestMin(t *testing.T) {
	array := []int{10, 15, 7, 40, 52, 12, 60, 777}
	expected_index := 2
	expected_value := 7
	received_index, received_value := min(array)
	if received_index != expected_index || received_value != expected_value {
		t.Errorf("Error, got: %d, %d, want: %d, %d", received_index, received_value, expected_index, expected_value)
	}
}
```
