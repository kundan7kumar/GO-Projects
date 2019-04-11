// "var" : declares 1 or more variables.Also declare multiple variables at once.
// It will infer the type of initialized variables.
// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.

package main

import "fmt"

func main() {

	var a = "India"
	fmt.Println(a)

	var b, c int = 5, 7
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}
