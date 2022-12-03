package nomath

import (
	"strconv"
	"strings"
)

func MeasureWrappingPaperSurface(input string) (surface int) {

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimPrefix(input, "\n")
	presentsList := strings.Split(input, "\n")

	for _, present := range presentsList {

		dimensionsList := strings.Split(present, "x")

		x, y, z := parseDimensions(dimensionsList)
		xy, yz, zx, surfaceOfThisPresent := calculateAllSurfaces(x, y, z)
		slackOfThisPresent := smallestSurface(xy, yz, zx)

		surface += surfaceOfThisPresent + slackOfThisPresent
	}

	return
}

func calculateAllSurfaces(x, y, z int) (xy int, yz int, zx int, total int) {

	xy = x * y
	yz = y * z
	zx = z * x

	return x * y, y * z, z * x, 2*xy + 2*yz + 2*zx
}

func parseDimensions(dimensionsList []string) (int, int, int) {
	xStr, yStr, zStr := dimensionsList[0], dimensionsList[1], dimensionsList[2]
	x, _ := strconv.ParseInt(xStr, 10, 64)
	y, _ := strconv.ParseInt(yStr, 10, 64)
	z, _ := strconv.ParseInt(zStr, 10, 64)
	return int(x), int(y), int(z)
}

func smallestSurface(xy, yz, zx int) int {
	if xy < yz {

		if xy < zx {
			return xy
		} else {
			return zx
		}
	} else {

		if yz < zx {
			return yz
		} else {
			return zx
		}
	}
}
