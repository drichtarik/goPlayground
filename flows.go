package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

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

func makeSqrt(input float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - input) / (2*z)
	}
	return z
}

// Switch is evaluated from top to bottom
func getOSName() {
	fmt.Print("This machine runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n" , os)
	}
}

// This feels unnatural
func whenDoIHaveToWork() {
	fmt.Print("When do I have to work? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("In two days.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Today.")
	}
}

func whenDoesTheweekendStart() {
	fmt.Print("When does the weekend start? ")
	today := time.Now().Weekday()
	switch today {
	case time.Monday:
		fmt.Println("In five days.")
	case time.Tuesday:
		fmt.Println("In four days.")
	case time.Wednesday:
		fmt.Println("In three days")
	case time.Thursday:
		fmt.Println("In two days.")
	case time.Friday:
		fmt.Println("Tomorrow")
	case time.Saturday:
		fmt.Println("Already started today")
	case time.Sunday:
		fmt.Println("Already started yesterday.")
	default:
		fmt.Println("Guess I'm broken :(")
	}
}

// switch with no condition
func whichPartOfDayIsNow()  {
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Good morning!")
	case time.Now().Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")

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

	var a, b, lim float64 = 3, 2 ,10
	// only if scoped variables
	if v := math.Pow(a, b); v < lim {
		fmt.Println(v)
	}
	fmt.Println(makeSqrt(36))
	getOSName()
	whenDoIHaveToWork()
	whenDoesTheweekendStart()
	whichPartOfDayIsNow()
}