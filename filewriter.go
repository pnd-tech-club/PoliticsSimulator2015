package main

import (
"fmt"
//"os" //for when file io is in here
)

/* Not to be called explicitly, run BuildPage first, will add import variables later */
func WriteOut() {
	bar := "+-------------------------------------------------------------------+"
	sd := "|"
	worden := "test, test" // temporary

	fmt.Println(bar)
	fmt.Printf("%v\n%v %s %v\n%v", bar, sd, worden, sd, bar)
	/* file io should replace prints */
}