package notquitelisp

import "strings"

func TranslateDirectionsToFloor(directions string) (floor int) {

	split := strings.Split(directions, "")

	for _, direction := range split {

		floor = getNextFloorFor(direction, floor)

	}

	return
}

func EntersTheBasementAt(directions string) (position int) {

	split := strings.Split(directions, "")

	var currentFloor int
	for _, direction := range split {

		position++

		currentFloor = getNextFloorFor(direction, currentFloor)

		if currentFloor < 0 {
			return
		}
	}

	return
}

func getNextFloorFor(direction string, currentFloor int) int {

	if direction == "(" {

		return currentFloor + 1
	}

	return currentFloor - 1
}
