package main

import (
"fmt"
"os"
)

/****
All functions contained herein should be solely called (optimally?) from the
main (main.go's main()) game loop's case structure used to interpret user input.
****/


/* should be able to load in a `.save` file and then restore/set
global variables/stats/states/stages via reading from the file */
func load(filename string) {
	fmt.Print("DEBUG: This function is still <WIP>. Filename: ", filename)
}


/* if whoops is true on return, it means there is an error */
func save()(whoops bool) {
	var (
	savename string
	safetosave bool
	)
	whoops = false
	safetosave = true

	fmt.Println("DEBUG: This function is not completed.")
	fmt.Print("Save-file name?: ")
	fmt.Scan(&savename)
	fmt.Printf("Writing %s.save  :)\n", savename)

	dir, err := os.Open(".")
	if err != nil {
         	fmt.Print("Bad, bad things happened here (1)")
		whoops = true
		return
         }
         defer dir.Close()

         fileInfos, err := dir.Readdir(-1)
         if err != nil {
         	fmt.Print("Bad, bad things happened here (2)")
		whoops = true
		return
         }

         for _, fi := range fileInfos {
         	fmt.Println("DEBUG: ", fi.Name())
		if fi.Name() == (savename + ".save") {
			safetosave = false
		}
         }
	if safetosave == false {
		fmt.Print("That file already exists! *bap*")
		whoops = true
		return
	} else {
		file, err := os.Create((savename + ".save"))
		if err != nil {
			fmt.Print("Failed to create file!")
			whoops = true
			return
		}
		defer file.Close() // defer executes when function "Returns"

		file.WriteString("test for savefile") // temporary write, eventually we will have a set of vals
	}
	return
}

func help() {
	fmt.Printf("\nCommands: \nhelp\nsave\nquit\n")
}
