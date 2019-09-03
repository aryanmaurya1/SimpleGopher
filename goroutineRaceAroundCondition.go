/*
This simple GO program demonstrate the race around
condition using goroutines.

This program has two functions 'goroutines1' and
'goroutines2' which try to modify a global variable
x which is an integer. Both functions are called as
goroutines inside main() function.

Order of execution of these functions is not defined,
there will a chance that both functions try to change
the value of variable x at same time which will lead to
a race around condition.

Finally program will crash or the system stops working.
*/

package main

import "fmt"

var x int

func goroutine1() {
	x++
	fmt.Println("BY GOROUTINE1 - ", x)
}

func goroutine2() {
	x--
	fmt.Println("BY GOROUTINE2 - ", x)
}

func main() {
	for {
		go goroutine1()
		go goroutine2()

	}

}
