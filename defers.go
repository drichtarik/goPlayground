package main

import (
	"fmt"
)

func basicDefer()  {
	defer fmt.Println("World!")
	fmt.Print("Hello ")
}

//  Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func stackedDefers() {
	fmt.Println("Let's count to 10 with defer")
	defer fmt.Println("This is when the function started.")
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done!")
}

func main() {
	basicDefer()
	stackedDefers()
}
