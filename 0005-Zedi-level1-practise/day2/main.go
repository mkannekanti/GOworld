package main

import "fmt"

// Day 2:
// 1. 	Use var to DECLARE three VARIABLES.
//		The variables should have package level scope.
//		Do not assign VALUES to the variables.
//		Use the following IDENTIFIERS for the variables and make sure
//		the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
// 			- identifier “x” type int
// 			- identifier “y” type string
// 			- identifier “z” type bool
// 2.	in func main
// 			- print out the values for each identifier
// 			- The compiler assigned values to the variables. What are these values called?

var x int
var y string
var z bool

func main() {
	fmt.Printf("Zero values of int is %#v, string is %#v, boolean is %#v \n", x, y, z)
}
