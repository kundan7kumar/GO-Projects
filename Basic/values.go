//Values type in GO Booleans.
//a. Strings, which can be added together with +.
//b. Integers and floats.
//c. with boolean operators
//Few basic examples.

package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("5+4 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
