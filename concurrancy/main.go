package main

import (
	"context"
	"fmt"
	"log"
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

func sender(ctx context.Context, intCh chan Message) {
	defer wg.Done()
	defer fmt.Println("gorutine send exit")
	fmt.Println("gorutine send")
	for {
		select {
		case <-ctx.Done():
			log.Fatal("Time is out")
			return
		case <-time.Tick(2 * time.Second):
			intCh <- Message{time.Now(), "time"}
		}
	}
}
func reciever(ctx context.Context, intCh chan Message) {
	defer wg.Done()
	defer fmt.Println("gorutine send exit")
	fmt.Println("gorutine rec")

	for {
		select {
		case <-ctx.Done():
			fmt.Print(ctx.Err())
			return
		case <-time.Tick(5 * time.Second):
			print(<-intCh)
		}
	}

}

func main() {
	wg.Add(2)
	intCh := make(chan Message)
	//defer close(intCh)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = ctx
	defer cancel()
	group := func() {
		//defer wg.Done()
		defer close(intCh)
		go sender(ctx, intCh)
		go reciever(ctx, intCh) //не заканчивается
		wg.Wait()
	}
	group()

	//for {
	//	select {
	//	case <-time.After(time.Microsecond):
	//		wg.Add(3)
	//		go group()
	//		wg.Wait()
	//	case <-ctx.Done():
	//		fmt.Println(ctx.Err())
	//		return
	//	}
	//}
}
