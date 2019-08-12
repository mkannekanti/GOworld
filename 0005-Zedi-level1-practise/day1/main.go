package main

import "fmt"

// Goal:
// 1. 	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with
// 		the IDENTIFIERS “x” and “y” and “z”
//		- 42
// 		- “James Bond”
// 		- true
// 2.	Now print the values stored in those variables using
//		- a single print statement
//		- multiple print statements

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("x:", x, ", y:", y, ", z:", z)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
