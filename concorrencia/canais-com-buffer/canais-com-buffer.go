package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)
	go testandoChan("oii", canal)
	msg, aberto := <-canal
	if !aberto {
		msg = "nao é possivel retornar o aberto"
	}
	msg2, aberto2 := <-canal
	if !aberto2 {
		msg = "nao é possivel retornar o aberto 2"
	}

	fmt.Println(msg)
	fmt.Println(msg2)

}

func testandoChan(texto string, ch chan string) {
	ch <- texto
	ch <- "teste"
	close(ch)
}
