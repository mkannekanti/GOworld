package main

import "fmt"

func main() {
	fmt.Println("Hello GO World!!")

	// control flow:
	// [1] sequence
	// [2] loop; iterative
	// [3] conditional

	// sequence
	fmt.Println("Testing out the Sequence..")
	foo()

	// loop
	fmt.Println("Testing out the loop and iteratives..")

	// Wait!! Wait!! Let me take a note here!!
	// 'i' is declared using ":="
	// no () for for and if
	// Way to Go
	for i := 0; i < 10; i++ {

		// conditional
		if i%2 == 0 {
			fmt.Println("Iterating on even numbers", i)
		}
	}
	fmt.Println("Loop, iteratives are looking good.")

	// sequence again
	bar()
}

func foo() {
	// Check how "_" is used to ignore the return value
	n, _ := fmt.Println("I'm in foo")
	fmt.Println("Sequence looks good. Here is your reward of", n, "GOld coins")
}

func bar() {
	n, err := fmt.Println("Hey GO world, everything is fine here. Let me stick around for sometime!!")
	fmt.Println("Last statement has", n, "GOld coins but their worth is", err, ":P")
}
