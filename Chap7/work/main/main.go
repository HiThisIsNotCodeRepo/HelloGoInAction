package main

import (
	"fmt"
	"github.com/qinchenfeng/HelloGoInAction/Chap7/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
	ack  chan struct{}
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(5 * time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(1 * len(names))

	for i := 0; i < 1; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
				fmt.Println("done")
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
