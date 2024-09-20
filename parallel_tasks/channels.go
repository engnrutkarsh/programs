package main

import (
	"fmt"
	"time"
)

func greet(text string, doneChan chan bool) {
	fmt.Println(text)
	doneChan <- true

}
func slowGreet(text string, doneChan chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println(text)
	doneChan <- true
}
func main() {
	done := make(chan bool)
	go greet("syste", done)
	go slowGreet("Hello", done)
	go greet("system", done)
	<-done
	<-done
	<-done
}
