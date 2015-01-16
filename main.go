/*
Test
*/

package main

import (
"fmt"
//"math"
//"math/cmplx"
//"io"
//"strings"
"svi"
)

var version = "0.1a"

func main() {

var usrrtn string
var randz int
tf := false

	fmt.Println("Welcome to PND Politics Simulator 2015")
	fmt.Println("Version: ", version)
	fmt.Println("Written and Built for your Pleasure by PND Tech Club")
	fmt.Println("Enjoyment is Optional.")
	usrrtn, tf = svi.YorN("Select Your Metatype?")
	randz = svi.Random(1, 6)
	fmt.Println(usrrtn, tf, randz) // testing
}

