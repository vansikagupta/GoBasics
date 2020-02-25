package main

import (
	"fmt"
	"time"
	"sync"
)

func cleanup() {
	//Handling the errpr using recover()
	if r := recover(); r != nil { 
		fmt.Println("Recovered in cleanup :",r)
	}
}

func foo(msg string) {
	defer wg.Done()
	defer cleanup()
	for i := 0; i < 5; i++{
		fmt.Println(msg)
		time.Sleep(100*time.Millisecond)
		if i == 2 {
			panic("Oh Dear its a 2") // creating an error
		}
	}
}


var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go foo("Hey")
	wg.Add(1)
	go foo("There")
	wg.Wait()
}