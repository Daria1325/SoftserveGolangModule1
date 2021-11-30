package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Message struct {
	time time.Time
	msg  string
}

func print(indiv Message) {
	fmt.Println(indiv.time.Format("15:04:05"), indiv.msg)
}

func sender(intCh chan Message, quit chan bool, quitrecv chan bool) {
	defer wg.Done()
	defer fmt.Println("\ngorutine send exit")

	for {
		select {
		case <-quit:
			quitrecv <- true
			return
		case <-time.Tick(2 * time.Second):
			intCh <- Message{time.Now(), "time"}
		}
	}
}
func reciever(intCh chan Message, quitrecv chan bool) {
	defer wg.Done()
	defer fmt.Println("\ngorutine rec exit")
	for {
		select {
		case <-quitrecv:
			for i := 0; i < len(intCh); i++ {
				print(<-intCh)
			}
			return
		case <-time.Tick(5 * time.Second):
			for i := 0; i < len(intCh); i++ {
				print(<-intCh)
			}
		}
	}

}

func main() {
	wg.Add(3)
	intCh := make(chan Message, 100)
	quit := make(chan bool)
	quitrecv := make(chan bool)

	defer close(intCh)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = ctx
	defer cancel()
	go func() {
		defer wg.Done()
		go sender(intCh, quit, quitrecv) //don't exit
		go reciever(intCh, quitrecv)

		for {
			select {
			case <-ctx.Done():
				fmt.Print(ctx.Err())
				quit <- true
				return
			default:

			}
		}
	}()
	wg.Wait()

}
