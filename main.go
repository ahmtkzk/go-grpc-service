package main

import (
	"fmt"
	"time"
)

var (
	ch = make(chan string)
)

func main() {

	go hello()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	time.Sleep(time.Second * 3)

}

func hello() {
	fmt.Println("Merhaba")
	ch <- "Selam"
	fmt.Println(<-ch)
	ch <- "Naber"
}
