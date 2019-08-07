package main

import "fmt"

// package scoped variables. Without initializer it's 0, false or ""
var c, python, java bool
// initialized variables
var q, w int = 1, 2
var e, word, statement = 3, "word", true

// constants use keyword const cannot be declared using ":="
const brno  = "Brno"

// simple function
func typeOut(x, y string) (string, string) {
	return x, y
}

// named return values
func namedReturnValues(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "naked" return returns the named values by default. Use only in short, simple functions
	return
}

func basicTypes() {
	fmt.Println("Types in GO:")
	fmt.Println("bool")
	fmt.Println("string")
	fmt.Println("int  int8  int16  int32  int64")
	fmt.Println("uint uint8 uint16 uint32 uint64 uintptr")
	fmt.Println("byte - alias for uint8")
	fmt.Println("rune - alias for int32")
	fmt.Println("float32 float64")
	fmt.Println("complex64 complex128")
}

func main() {
	fmt.Println(typeOut("Hello", "World!"))
	fmt.Println(namedReturnValues(5))
	fmt.Println(brno)
	// function scoped variable
	var i int
	fmt.Println(c, python, java, i, 1, 2, 3, word, statement)
	// only in functions the "var" can be omitted with ":="
	r := 5
	fmt.Println(r)
	basicTypes()
	// to print type and value of a variable use fmt.Printf with %v for value and %T for type
	fmt.Printf("Value: %v is of type: %T\n", r, r)
	// type conversions
	var t int = 5
	u := float64(t)
	fmt.Printf("Value: %v is of type: %T and was %T\n", u, u, t)
}