package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
		time.Sleep(200* time.Millisecond)
	}
}