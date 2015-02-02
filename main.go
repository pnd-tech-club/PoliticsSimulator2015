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
	usrinpt string
)

tf := false


	fmt.Println("Welcome to PND Politics Simulator 2015")
	fmt.Println("Version: ", version)
	fmt.Println("Written and Built for your Pleasure by PND Tech Club")
	fmt.Println("Enjoyment is Optional.")
	
	// Basic Yes or No Dialogue, see `svi.go` for svi.*  functions
	for {
		usrrtn, tf = svi.YorN("Select Your Metatype?") // no [y/n]: is needed, as it is added
		if tf == true {
			fmt.Println("Okay.")
			break
		} else {
			fmt.Println("Pick - yes or no - not that hard")
		}
	}

	//randz := svi.Random(1, 6) // gen a number 1-5 (arbitrary)
	
	// Ask for metatype, then read from the relevant .meta file
	if usrrtn == "y" {
		for {
			fmt.Print("Metatype [? for list]: ")
			fmt.Scan(&usrinpt)
			if usrinpt == "?" {
				// There MUST be a better way to do this ... then again ...
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



	fmt.Println("Thanks for Playing :) ")
} // Program Ends

