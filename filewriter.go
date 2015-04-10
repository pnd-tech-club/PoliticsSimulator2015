package main

import (
"fmt"
/* gofpdf: https://github.com/jung-kurt/gofpdf */
"github.com/jung-kurt/gofpdf"
//"os" //for when file io is in here
"math"
)

var bar string = "+-------------------------------------------------------------------+"
var sd string = "|"
var id string = "Politics Simulator 2015 v" + version

/* Not to be called explicitly, run BuildPage first, will add import variables later */
func writeout() {
	worden := "test, test" // temporary

	fmt.Println(bar)
	fmt.Printf("%v\n%v %s %v\n%v", bar, sd, worden, sd, bar)
	/* file io should replace prints */
}

func printpdf(textbox string, filename string)(succ bool) {
	var x, y float64
	/* Portrait, inches, Letter, irrelevant fontdir (unless custom font) */
	pdf := gofpdf.New("P", "in", "Letter", "../font")
	pdf.AddPage()
	/* Arial font, no B/I/U formatting, size 14 */
	pdf.SetFont("Arial", "", 10)
	/* set/get position */
	x, y = pdf.GetXY()
	if x != 0.0 && y != 0.0 {
		pdf.SetXY(0.0, 0.0)
	}
	/* Cell for text, 7"x11", with text */
	//pdf.Cell(1, 1, bar)
	/* write out at current (X,Y) w/ line height 0.125" (as per .New())*/
	y = y+(0.1*(math.Pow(10, -1000)))
	pdf.SetY(y)
	pdf.Write(0.031250, id)
	/* .Write() breaks on newline, []string is thus unnecessary */
	pdf.SetY(y+1)
	pdf.Write(0.031250, textbox)
	//pdf.Output(os.Stdout) //succ relies on successful formatting/composition
	pdf.OutputFileAndClose(filename)

	succ = true //temporary until debugging implemented
	return
}
