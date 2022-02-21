package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func producers(c chan int, name string, maxInt int) {
	for i := 0; i < maxInt; i++ {
		fmt.Printf("producer name: %s, send: %d\n", name, i)
		c <- i
	}
}

func consumer(c chan int, name string, notifyC chan<- struct{}) {
	for {
		i, ok := <-c
		if ok {
			fmt.Printf("    consumer name: %s, recv: %v\n", name, i)
			notifyC <- struct{}{}
		} else {
			break
		}
	}
}

func main() {
	msgChannel := make(chan int, 100)
	notifyChannel := make(chan struct{})
	sum := 0
	consumerIndex := 0
	for proIndex := 0; proIndex < 10; proIndex++ {
		msgNums := rand.Intn(10)
		sum += msgNums
		go producers(msgChannel, strconv.Itoa(proIndex), msgNums)
		if rand.Intn(20)%3 == 0 {
			go consumer(msgChannel, strconv.Itoa(consumerIndex), notifyChannel)
			consumerIndex++
		}
	}

	for {
		<-notifyChannel
		sum -= 1
		if sum == 0 {
			fmt.Printf("all consumer over\n")
			close(notifyChannel)
			close(msgChannel)
			break
		}
	}
}
