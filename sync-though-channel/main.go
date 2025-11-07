package main

import "fmt"

var shared int

func main() {
	ch := make(chan struct{})

	go func() {
		ch <- struct{}{}
		shared = 10
	}()

	<-ch
	fmt.Println(shared)
}
