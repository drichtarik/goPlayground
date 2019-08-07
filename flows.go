package main

import "fmt"

// for loop with all the components
func showMeSomeNumbers() {
	sum := 1
	for i := 1; i <= 10; i++ {
		fmt.Println(i, sum)
		sum += i
	}
}

// init and post parameters are optional
func doItLessVerbose() {
	sum := 1
	for ; sum < 60; {
		fmt.Println(sum)
		sum += sum
	}
}

// now drop the semicolons and it's like a while loop
func forIsLikeWhile() {
	sum := 1
	for sum < 60 {
		fmt.Println(sum)
		sum += sum
	}
}

func main() {
	showMeSomeNumbers()
	doItLessVerbose()
	forIsLikeWhile()

	// () parentheses in if function are optional
	if true == true {
		fmt.Println("What did you expect, ha?")
	}
}