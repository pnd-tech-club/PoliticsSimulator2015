package main

import (
"io/ioutil"
"strings"
)

/* Pstats moved to engine.go */


func Metareader(metaname string)(newpstats Pstats, success int) {
	content, err := ioutil.ReadFile(metaname)
	if err == nil {

		lines := strings.Split(string(content), "\n")
		newpstats = Pstats{lines[0], lines[1], lines[2], lines[3], lines[4], lines[5], lines[6], lines[7], lines[8]}

		success = 0 //success
	} else {
		success = 1 //failure
		var lines [2]string
		lines[0] = "nil"
	}
	return
}
