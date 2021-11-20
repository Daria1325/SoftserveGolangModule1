package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Message struct {
	time time.Time
	msg  string
}

func print(indiv Message) {
	fmt.Println(indiv.time.Format("15:04:05"), indiv.msg)
}

func main() {
	var wg sync.WaitGroup

	intCh := make(chan Message)
	send := func(intCh chan Message) {
		defer wg.Done()
		intCh <- Message{time.Now(), "time"}
		time.Sleep(2 * time.Second)
	}
	recv := func(intCh chan Message) {
		defer wg.Done()
		i := <-intCh
		print(i)
		time.Sleep(5 * time.Second)
	}
	group := func() {
		defer wg.Done()
		go send(intCh)
		go recv(intCh)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	_ = ctx
	defer cancel()

	for {
		select {
		case <-time.After(time.Microsecond):
			wg.Add(3)
			go group()
			wg.Wait()
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
