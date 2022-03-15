package main

import (
	"fmt"
	"time"
)

func ticker(inc int32, done chan bool) {

	go func() { //inc2 int32
		//inc3 := inc2
		ticker := time.NewTicker(200 * time.Millisecond)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("inc: %v \n", inc)
			}
		}
		/*for tick := 1; tick <= 5; tick++ {
			fmt.Printf("Текущее значение переменной: %v\n", inc)
			time.Sleep(200 * time.Millisecond)
		}*/
	}()
}

func timer(s int32) {
	done := make(chan bool)
	var inc int32
	inc = 0
	for {
		if s <= 0 {
			break
		} else {
			ticker(inc, done)
			time.Sleep(1 * time.Second)
			s--
			//atomic.AddInt32(&inc,1)
			inc++
			done <- true

		}
	}
}

func main() {

	var a int32
	fmt.Println("Введите количество секунд: ")
	fmt.Scan(&a)
	timer(a)
}
