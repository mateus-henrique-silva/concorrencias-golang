package main

import (
	"fmt"
	"sync"
	"time"
)

// chamar um funcao e executar sua funcionalidade sem esperar que um outra se encerre
func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	go func() {
		escrever("Ola mundo")
		waitgroup.Done()
	}()
	go func() {
		escrever("teste golag")
		waitgroup.Done()
	}()
	waitgroup.Wait()
}
func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
