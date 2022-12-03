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

func EntersTheBasementAt(directions string) (position int) {
	split := strings.Split(directions, "")

	var currentFloor int
	for _, direction := range split {

		position++

		if direction == "(" {

			currentFloor += 1
		} else {

			currentFloor -= 1
		}

		if currentFloor < 0 {
			return
		}
	}

	return
}
