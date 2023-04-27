package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.NewTicker(2 * time.Second)
	for i := 0; i < 10; i++ {
		t := <-timer.C
		fmt.Println(t.Sub(start).Seconds())
	}
	timer.Stop()
}
