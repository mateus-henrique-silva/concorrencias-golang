package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "teste de canal 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal2 <- "teste de canal 2"
		}
	}()

	for {
		select {
		case menssagem1 := <-canal1:
			fmt.Println(menssagem1)
		case menssagem2 := <-canal2:
			fmt.Println(menssagem2)
		}
	}
}
