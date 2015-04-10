/*
PND Politics Simulator 2015
Conceived by Sean Hinchee

Disclaimer: Any and all relation or implication that may or may not relate to people who may or may not exist is entirely coincidence and a figment within the beholder's imagination. Please enjoy PS2015 responsibly.
*/

package main

import (
"fmt"
"svi"
"flag"
)

var (
version = "0.1a"
plyr Pstats
sucerr int
metapath string
)


func main() {
	flag.StringVar(&metapath, "f", "./", "Path to folder containing .meta files")
	flag.Parse()

	var (
		usrrtn string
		usrinpt string
		debug bool
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
			for  {
				// rewrite nested "for's" using continue's and no break's
				for {
					fmt.Print("Metatype [? for list]: ")
					fmt.Scan(&usrinpt)
					if usrinpt == "?" {
						// There MUST be a better way to do this ... then again ...
						fmt.Printf("\nInstigator - troublemaking bastards\nPreppy - \nStudent Ambassador - \nUnifier - \n\n")
					}
					metaname := usrinpt + ".meta"
					plyr, sucerr = Metareader(metaname)
					if sucerr == 0 {
						break
					}
				}
				fmt.Printf("\nYour stats: \n")
				fmt.Printf("Class: %s\nPower: %s\nInfluence: %s\nIntelligence: %s\nWealth: %s\nReputation: %s\nBlood Type: %s\nPolitical Spectrum Rating: %s\nHouse: %s\n\n", plyr.name, plyr.pwr, plyr.inf, plyr.int, plyr.wlth, plyr.rpt, plyr.bt, plyr.pol, plyr.hou)
				if sucerr == 0 {
					break
				}
			}
		} else {
			fmt.Println("Using the Default Metatype.")
			plyr, sucerr = Metareader("Default.meta")
			fmt.Printf("\nYour stats: \n")
			fmt.Printf("Class: %s\nPower: %s\nInfluence: %s\nIntelligence: %s\nWealth: %s\nReputation: %s\nBlood Type: %s\nPolitical Spectrum Rating: %s\nHouse: %s\n\n", plyr.name, plyr.pwr, plyr.inf, plyr.int, plyr.wlth, plyr.rpt, plyr.bt, plyr.pol, plyr.hou)
			if sucerr == 1 {
				fmt.Println("Could not read the 'Default.meta' file.")
			}
		}


		/* Begin main game loop */


		fmt.Println("Type '?' or 'help' to display a list of commands.")

		debug = false
		for running := 0; running < 1; {
			var response string
			fmt.Printf("\n%v> ", plyr.name)
			//fmt.Scan(&response)
			fmt.Scanln(&response)

			if debug == true {
				fmt.Printf("\nDEBUG: %v\n", response)
			}

			// switch or other processing locally, then handing off to cmds.go?
			switch response {
				case "exit":
					running = 1
				case "quit":
					running = 1
				case "help":
					help()
				case "?":
					help()
				case "save":
					save()
				case "debug":
					debug = debugging(debug)
				case "print":
					// printpdf() is in filewriter.go; prepprint() in cmds.go
					prepprint()
				case "":
					continue
				default:
					fmt.Println("Command not found.")
			}

		}

		fmt.Printf("\n\nThanks for Playing :) \n")
} // Program Ends
