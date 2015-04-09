package main

import (
"fmt"
/* gofpdf: https://github.com/jung-kurt/gofpdf */
"github.com/jung-kurt/gofpdf"
//"os" //for when file io is in here
)

var bar string = "+-------------------------------------------------------------------+"
var sd string = "|"

/* Not to be called explicitly, run BuildPage first, will add import variables later */
func writeout() {
	worden := "test, test" // temporary

	fmt.Println(bar)
	fmt.Printf("%v\n%v %s %v\n%v", bar, sd, worden, sd, bar)
	/* file io should replace prints */
}

func framepdf(textbox []string, filename string)(succ bool) {

	/* Portrait, inches, Letter, irrelevant fontdir (unless custom font) */
	pdf := gofpdf.New("P", "in", "Letter", "../font")
	pdf.AddPage()
	/* Arial font, no B/I/U formatting, size 14 */
	pdf.SetFont("Arial", "", 14)
	/* Cell for text, 7"x11", with text */
	pdf.Cell(20, 5, bar)
	//pdf.Output(os.Stdout) //succ relies on successful formatting/composition
	pdf.OutputFileAndClose(filename)

	succ = true //temporary until debugging implemented
	return
}
