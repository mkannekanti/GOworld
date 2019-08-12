package main

import "fmt"

// Regular Style
var coin int

// Rajinikath iStyle
type goldcoin int

var gcoin goldcoin

func main() {
	// Assigning values to Package level variables
	coin = 007  // Bond, James Bond.
	gcoin = 700 // Star, Super Start. Mind it Rascala!!

	fmt.Println("Far Far away, in the GO World, there are two heros, ")
	fmt.Println("Mr. James Bond and Mr. Rajinikanth a.k.a Thalaiva.")

	fmt.Println("")

	fmt.Println("Mr.Bond's Coin is:", coin)
	fmt.Println("Sorry!! We hurt Mr. Bond. Let's print his coin properly.")

	fmt.Println("")

	fmt.Printf("Mr. Bond's Coin is: %03d", coin)
	fmt.Printf(" and it's type is: %T \n", coin)
	fmt.Print("Mr. Rajini's Coin is: ", gcoin)
	fmt.Printf(" and it's type is: %T\n", gcoin)

	fmt.Println("")
	fmt.Println("One fine day, after Rajini saved the GO world 314159th time...")
	fmt.Println("")

	fmt.Println("Bond to Rajini:")
	fmt.Println("Thalaiva, I want to become like you. I want to get a GOld coin just like you.")
	fmt.Println("How can I do that?")

	fmt.Println("")

	fmt.Println("Rajini to Bond:")
	fmt.Print("My 'type' is seperate. ")
	fmt.Print("Better don't try to follow me. ")
	fmt.Println("Mind it!!")

	fmt.Println("")

	fmt.Println("Bond to Rajini:")
	fmt.Println("Thalaiva, please Thalaiva!!")

	fmt.Println("")

	fmt.Println("Rajini to Bond:")
	fmt.Println("Okay Macha. Just this once. You have to use this amazing GO conversion cape.")
	fmt.Println("Take my goldcoin(x) cape")

	fmt.Println("")

	fmt.Println("Bond to Rajini:")
	fmt.Println("Thank you boss!!!")

	var bondGcoin goldcoin = goldcoin(coin)

	fmt.Println("")

	fmt.Println("Mr.Bond to 'GO World':")
	fmt.Printf("Hurray!!! My type is now: %T \n", bondGcoin)
}
