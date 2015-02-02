package main

import (
"fmt"
"io/ioutil"
"strings"
)

type Pstats struct{
	name string
	pwr string
	inf string
	int string
	wlth string
	rpt string
	bt string
}


func Metareader(metaname string)(newpstats Pstats, success int) {
	content, err := ioutil.ReadFile(metaname)
	if err == nil {
		
		lines := strings.Split(string(content), "\n")
		newpstats = Pstats{lines[0], lines[1], lines[2], lines[3], lines[4], lines[5], lines[6]}

		success = 0 //success
	} else {
		fmt.Println("That's not a valid Metatype.")
		success = 1 //failure
		var lines []string
		lines[0] = "nil"
	}
	return
}
