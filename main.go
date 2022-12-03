package main

import (
	"fmt"
	"os"

	notquitelisp "github.com/ppicom/past-advents-of-code/2015/1-not-quite-lisp"
	nomath "github.com/ppicom/past-advents-of-code/2015/2-no-math"
)

func main() {
	notquitelispdata, err := os.ReadFile("input_notquitelisp.txt")
	if err != nil {
		panic(err)
	}

	floor := notquitelisp.TranslateDirectionsToFloor(string(notquitelispdata))
	fmt.Printf("Santa's floor is: %d\n", floor)

	position := notquitelisp.EntersTheBasementAt(string(notquitelispdata))
	fmt.Printf("Santa enters the basement at: %d\n", position)

	nomathdata, err := os.ReadFile("input_nomath.txt")
	if err != nil {
		panic(err)
	}

	surface := nomath.MeasureWrappingPaperSurface(string(nomathdata))
	fmt.Printf("Wrapping paper surface: %d\n", surface)

	ribbon := nomath.MeasureRibbonLength(string(nomathdata))
	fmt.Printf("Ribbon length: %d\n", ribbon)
}
