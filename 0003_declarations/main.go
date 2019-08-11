package main

import "fmt"

// := is Short declaration of a variable
// "var" keyword also declares a variable
//
//	|	Short Declaration	|	var Keyword			|
//	|-----------------------------------------------|
//	| Must be inside a 		| It can be global.		|
//	| function body( {} )	| Used for accessing	|
//	|						| throughout the file	|
//	|-----------------------------------------------|
//

var y = 10
var z int
var z0 int = 100

func main() {
	x := "C"
	g1, _ := fmt.Println("Hey!! You are", x, ". You have", y, "GOld coins")

	g2, _ := fmt.Println("Let me help you with some magic. Abrakadabra!! Abrakadabra!! BOOM!!")

	x = "Gopher"
	g3, _ := fmt.Println("There you go. You are now", x)

	y := y + g1 + g2 + g3 + 10
	fmt.Println("You now have", y, "GOold coins!! Enjoy my 10 GOld coins bonus")

	z = 50
	fmt.Println("Keeping", z, "Gold coins into Safe box")

	foo()
}

func foo() {
	t := z + z0
	fmt.Println("Total worth in the safe box is:", t)
}
