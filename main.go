package main

import (
	"fmt"
)
import "time"

func ticker(inc int) {
	go func() {
		for tick := 1; tick <= 5; tick++ {
			fmt.Printf("Текущее значение переменной: %v\n", inc)
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func timer(s int) {
	inc := 0
	for {
		if s <= 0 {
			break
		} else {
			ticker(inc)
			time.Sleep(1 * time.Second)
			s--
			inc++
		}
	}
}

func main() {

	var a int
	fmt.Println("Введите количество секунд: ")
	fmt.Scan(&a)
	timer(a)
}
