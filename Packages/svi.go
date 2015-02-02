// Package SVI for PNDTC
package svi

import (
"strings"
"fmt"
"math/rand"
"time"
//"bufio"
//"io"
"io/ioutil"
"os"
)


func filecheck(e error) {
	if e != nil {
		panic(e)
	}
}

func Metareader(metaname string)(success int) {
	content, err := ioutil.ReadFile(metaname)
	if err == nil {
		lines := strings.Split(string(content), "\n")
		fmt.Println(lines[0]) // first line
		fmt.Println(lines[len(lines)-1]) // last line
		success = 0 //success
	} else {
		fmt.Println("That's not a valid Metatype.")
		success = 1 //failure
	}
	return
}

// forgot I made this one, whoops
func metaread() {
	
	dat, err := ioutil.ReadFile("./instigator.meta")
	filecheck(err)
	fmt.Print(string(dat))
	
	buf := make([]byte, 1024)

	instigator, err := os.Open("./instigator.meta")
	filecheck(err)

	for {
		n, err := instigator.Read(buf)
		filecheck(err)
		if n == 0 {
			break
		}
	}
	
}


/* processes [y/n] dialogues with the yes/no user input
the values returned are the "y" or "n" and then a true/false for looping
the false value is stated false if the input is not a yes or no */

func YorN(prompt string) (newinputz string, tf bool) {
	inputz := "Mountain Goat"
	tf = false

	for tf != true {
	fmt.Printf("\n%v [y/n]: ", prompt)
	fmt.Scan(&inputz)
	
	if inputz = strings.ToLower(inputz); inputz == "yes" {
		newinputz = "y"
		tf = true
	} else if inputz == "no" {
		newinputz = "n"
		tf = true
	} else if inputz == "y" {
		tf = true
		newinputz = inputz
	} else if inputz == "n" {
		tf = true
		newinputz = inputz
	} else {
		newinputz = "Mountain Goat"
		tf = false
	} 

	}
	return
}


func Random(min, max int)(newnum int) {
rand.Seed(time.Now().Unix())
newnum = rand.Intn(max - min) + min //non-inclusive on max
return
}
