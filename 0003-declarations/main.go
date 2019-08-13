package main

import "fmt"

// := is Short declaration of a variable
// "var" keyword also declares a variable
//
//	|	Short Declaration	|	var Keyword			|
//	|-----------------------------------------------|
//	| Must be inside a 		| Package Level Scope.	|
//	| function body( {} )	| Used for accessing	|
//	|						| throughout the file	|
//	|-----------------------------------------------|
//

// Package Level scope variable
var y = 2
var safe = 0

// Zero Value of int is assigned to Z i.e., 0
var z int

// same as 'var z0 string = "100"'
// Zero value of String is an empty string
var z0 = "100"

// Boolean declaration with Zero value of "false"
var isSafe bool

func main() {

	// Short declaration
	x := "C"
	g1, _ := fmt.Println("Hey!! You are '", x, "'. Looks like you have only", y, "GOld coins")
	g2, _ := fmt.Println("Let me help you with some magic. Abrakadabra!! Abrakadabra!! BOOM!!")

	// Assigning a new value to a Short declared variable
	x = "GOpher(:=)"
	g3, _ := fmt.Println("There you go. You are now a", x)

	// Assigning 'var' declared variable with an expression
	z = y + g1 + g2 + g3 + 10
	fmt.Println("You now have", z, "GOold coins!! Enjoy my 10 GOld coins bonus")

	// Zero value is false
	if !isSafe {
		fmt.Println("Caution!! Your GOld coins are not safe in this world. Try to put them in safe.")
		isSafe = safebox()

		if isSafe {
			fmt.Println("Congratulations on keeping your GOld coins in the Safebox.")
		}
	}
}

// function "safebox()" takes no arguments and returns boolean
func safebox() bool {
	// Accesing Package Level var variables in other functions
	safe = z

	// C-like printing variables.
	fmt.Printf("These %#v GOld coins are of brand '%T'. Let's add them to your Safebox \n", safe, safe)
	fmt.Printf("Yay!! Amazon just delivered %#v GOld coins of branch '%T'\n", z0, z0)

	fmt.Println("Sorry! Our safebox currently takes only 'int' brand. Throw them out of GO World.")
	fmt.Println("Total worth in your Safebox is now:", safe)

	return true
}
