package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola mundo", canal)
	msg := <-canal
	fmt.Println(msg)
}
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
}
