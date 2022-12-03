package nomath

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func MeasureWrappingPaperSurface(input string) (surface int) {

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimPrefix(input, "\n")
	presentsList := strings.Split(input, "\n")

	for _, present := range presentsList {

		dimensionsList := strings.Split(present, "x")

		x, y, z := parseDimensions(dimensionsList)
		present := newPresent(x, y, z)

		surface += present.totalArea() + present.smallestArea()
	}

	return
}

func MeasureRibbonLength(input string) (ribbon int) {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimPrefix(input, "\n")
	presentsList := strings.Split(input, "\n")

	for _, present := range presentsList {

		dimensionsList := strings.Split(present, "x")

		x, y, z := parseDimensions(dimensionsList)
		present := newPresent(x, y, z)

		ribbon += present.shortestPerimeter() + present.volume()
	}

	return
}

func newPresent(x, y, z int) present {
	sides := []int{x, y, z}
	slices.Sort(sides)
	return present{
		sides: sides,
	}
}

type present struct {
	sides []int
}

func (p present) totalArea() int {

	xy := p.sides[0] * p.sides[1]
	yz := p.sides[1] * p.sides[2]
	zx := p.sides[2] * p.sides[0]

	return 2*xy + 2*yz + 2*zx
}

func (p present) smallestArea() int {

	areas := p.areasInIncreasingOrder()
	return areas[0]
}

func (p present) areasInIncreasingOrder() []int {

	x, y, z := p.sides[0], p.sides[1], p.sides[2]
	areas := []int{x * y, x * z, y * z}
	slices.Sort(areas)

	return areas
}

func (p present) shortestPerimeter() int {

	return 2*p.sides[0] + 2*p.sides[1]
}

func (p present) volume() int {
	x, y, z := p.sides[0], p.sides[1], p.sides[2]

	return x * y * z
}

func parseDimensions(dimensionsList []string) (int, int, int) {
	xStr, yStr, zStr := dimensionsList[0], dimensionsList[1], dimensionsList[2]
	x, _ := strconv.ParseInt(xStr, 10, 64)
	y, _ := strconv.ParseInt(yStr, 10, 64)
	z, _ := strconv.ParseInt(zStr, 10, 64)
	return int(x), int(y), int(z)
}
