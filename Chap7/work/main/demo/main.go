package main

import "time"

func main() {
	newChan := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		newChan <- struct{}{}
	}()
	<-newChan
}
