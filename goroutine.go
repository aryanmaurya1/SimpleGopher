package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `* + - . ! @ # $ % ^ & * ( )` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fmt.Println(fibo(45))
}
