package main

import (
	"fmt"
	"time"
)

func main() {

	chanInt := make(chan int, 10)

	go func() {
		tickerSend := time.NewTicker(1000 * time.Millisecond)
		for i := 0; ; i++ {
			<-tickerSend.C
			chanInt <- i
		}
	}()

	tickerRecv := time.NewTicker(1 * time.Second)
	for {
		<-tickerRecv.C
		ret := <-chanInt
		fmt.Printf("recv : %d\n", ret)
	}

}
