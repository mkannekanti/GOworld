package main

import (
	"fmt"
	"runtime"
)

// Numeric Types
// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32

var u8 uint8
var u16 uint16
var u32 uint32
var u64 uint64

var d8 int8
var d16 int16
var d32 int32
var d64 int64

var f32 float32
var f64 float64

var c64 complex64
var c128 complex128

var b byte
var r rune

func main() {
	fmt.Println("Welcome to the tour of GO World!!")
	fmt.Println("All the buildings here are built by", runtime.GOOS)
	fmt.Println("All of them are architected by ", runtime.GOARCH)

	fmt.Println("")

	fmt.Println("Guide: Here is puzzle to all!! Can you find the meaning of the writing on this building?")

	s := `
	uint8       the set of all unsigned  8-bit integers (0 to 255)
	uint16      the set of all unsigned 16-bit integers (0 to 65535)
	uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	
	int8        the set of all signed  8-bit integers (-128 to 127)
	int16       the set of all signed 16-bit integers (-32768 to 32767)
	int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	
	float32     the set of all IEEE-754 32-bit floating-point numbers
	float64     the set of all IEEE-754 64-bit floating-point numbers
	
	complex64   the set of all complex numbers with float32 real and imaginary parts
	complex128  the set of all complex numbers with float64 real and imaginary parts
	
	byte        alias for uint8
	rune        alias for int32
	
	uint     either 32 or 64 bits
	int      same size as uint
	uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value`

	// converting to slice of bytes
	bs := []byte(s)

	fmt.Println("..................................................")
	for i, v := range s {
		if i < 15 {
			fmt.Printf("%d+d+%d+%U+", i, bs[i], v)
		} else {
			break
		}
	}
	fmt.Println(" ....")
	fmt.Println("..................................................")

	fmt.Println("")

	fmt.Println("Person A: This looks hard man!! Must be prepared by Thalaiva. Any clue?")
	fmt.Println("Guide: Your guess is right!! This is indeed created by Mr. Rajini")
	fmt.Println("Here is the clue:")

	bbs := fmt.Sprintf("%T+%T", s, bs)
	for i := 0; i < len(bbs); i++ {
		fmt.Printf("%#X+", bbs[i])
	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("After two long hours....!!")
	fmt.Println("Person 314: I think I got it. Looks like it's a string that is sliced into bytes.")
	fmt.Println("And each byte then converted each byte to UTF-8 format. Even the clue is done that way.")
	fmt.Println("Guide: Amazing!! you're are on the right path!! Can you decode the full writing?")
	fmt.Println("Person 314: I think I can. Let me work on that")

	fmt.Println("")

	fmt.Println("After sometime....!!")
	fmt.Println("............................................................")
	fmt.Println(s)
	fmt.Println("............................................................")

	fmt.Println("")

	fmt.Println("Guide: Great!! Here is your gift of Rajinikanth's super convertion cape.")
	fmt.Println("Person 314: Wow!! Thanks Guide!! Thanks Thalaiva!! Thanks GO!!")
}
