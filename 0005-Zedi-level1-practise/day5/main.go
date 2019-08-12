package main

import "fmt"

// Goal:
// Addition to Day#4
// 1. 	at the package level scope, using the “var” keyword, create a VARIABLE
//		with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE
//		of your custom TYPE “x”
// eg:
//		type hotdog int
//		var x hotdog
//		var y int
//
// 2.	in func main
// 		-	this should already be done
// 			- print out the value of the variable “x”
// 			- print out the type of the variable “x”
// 			- assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// 			- print out the value of the variable “x”
// 3. 	now do this
// 		- now use CONVERSION to convert the TYPE of the VALUE stored in “x” to
//		  the UNDERLYING TYPE
// 			- then use the “=” operator to ASSIGN that value to “y”
// 			- print out the value stored in “y”
// 			- print out the type of “y”

type goldCoin int

var x goldCoin
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
