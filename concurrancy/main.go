//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//type Message struct {
//	time time.Time
//	msg string
//}
//func print(indiv Message){
//	fmt.Println(indiv.time.Format("15:04:05"),indiv.msg)
//}
//
//
//func main(){
//	var wg sync.WaitGroup
//	wg.Add(2)
//	intCh := make(chan Message)
//	send := func(intCh chan Message) {
//		defer wg.Done()
//		for{
//			intCh <- Message{time.Now(),"time"}
//			time.Sleep(2 * time.Second)
//		}
//	}
//	recv := func(intCh chan Message) {
//		defer wg.Done()
//		for{
//			i:=<-intCh
//			print(i)
//			time.Sleep(5 * time.Second)
//		}
//	}
//
//		go send(intCh)
//		go recv(intCh)
//		wg.Wait()
//}
//TODO context
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
	wg.Add(3)
	intCh := make(chan Message)
	send := func(intCh chan Message) {
		defer wg.Done()
		for {
			intCh <- Message{time.Now(), "time"}
			time.Sleep(2 * time.Second)
		}
	}
	recv := func(intCh chan Message) {
		defer wg.Done()
		for {
			i := <-intCh
			print(i)
			time.Sleep(5 * time.Second)
		}
	}
	group := func() {
		defer wg.Done()
		go send(intCh)
		go recv(intCh)
		wg.Wait()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	select {
	case <-time.After(time.Microsecond):
		go group()
		wg.Wait()
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		//return

		//wg.Wait()

	}
}
