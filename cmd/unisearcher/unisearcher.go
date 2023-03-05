package main

import (
	endpoints "assignment1/internal"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	go endpoints.StartServer(wg)

	wg.Wait()
}
