package main

import (
	"fmt"
	"time"
)

// chamar um funcao e executar sua funcionalidade sem esperar que um outra se encerre
func main() {
	go escrever("Ola mundo")
	escrever("teste golag")

}
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
