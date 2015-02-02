/*
Test
*/

package main

import (
"fmt"
//"math"
//"math/cmplx"
//"io/ioutil"
//"strings"
"svi"
)

var version = "0.1a"

func main() {

var (
	usrrtn string
	randz int
	usrinpt string
)

tf := false


	fmt.Println("Welcome to PND Politics Simulator 2015")
	fmt.Println("Version: ", version)
	fmt.Println("Written and Built for your Pleasure by PND Tech Club")
	fmt.Println("Enjoyment is Optional.")
	usrrtn, tf = svi.YorN("Select Your Metatype?") // no [y/n]: is needed, as it is added
	randz = svi.Random(1, 6) // gen a number 1-5
	
	if usrrtn == "y" {
		for {
			fmt.Print("Metatype [? for list]: ")
			fmt.Scan(&usrinpt)
			if usrinpt == "?" {
				fmt.Printf("\nInstigator - troublemaking bastards\nPreppy - \nStudent Ambassador - \nUnifier - \n")
			}
			metaname := usrinpt + ".meta"
			err := svi.Metareader(metaname)
			if err == 0 {
				break
			}
		}
	} else {
		fmt.Println("Using the Default Metatype")
		err := svi.Metareader("Default.meta")
		if err == 0 {
			fmt.Println("Could not read the 'Default.meta' file.")
		}
	}

	fmt.Println(usrrtn, tf, randz) // testing
}

