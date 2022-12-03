package notquitelisp

import "strings"

func TranslateDirectionsToFloor(directions string) (floor int) {
	split := strings.Split(directions, "")

	for _, direction := range split {

		if direction == "(" {

			floor += 1
		} else {

			floor -= 1
		}

	}

	return
}
