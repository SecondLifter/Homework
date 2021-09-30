package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	buf =make(chan int,10)
)

func Write(max int)  {
	for i:=0;i <= max;i++{
		buf <- i
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func Read()  {
	for {
		fmt.Println(<-buf)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func testWirteAndRead(max int)  {
	//wg.Add(1)
	go Write(max)
	//wg.Add(1)
	go Read()
}

func main()  {
	wg.Add(1)
	testWirteAndRead(10)
	wg.Wait()
}