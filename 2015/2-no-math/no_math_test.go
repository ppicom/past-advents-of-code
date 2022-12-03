package nomath

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {}

func (suite *TestSuite) TestWrappingPaperSurface() {

	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "2x3x4",
			output: 58,
		},
		{
			input:  "1x1x10",
			output: 43,
		},
		{
			input: `1x1x10
2x3x4`,
			output: 101,
		},
	}

	for _, tt := range tests {

		floor := MeasureWrappingPaperSurface(tt.input)
		suite.Equal(tt.output, floor)
	}
}
func (suite *TestSuite) TestRibbonLength() {

	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "2x3x4",
			output: 34,
		},
		{
			input:  "1x1x10",
			output: 14,
		},
		{
			input: `1x1x10
2x3x4`,
			output: 48,
		},
	}

	for _, tt := range tests {

		floor := MeasureRibbonLength(tt.input)
		suite.Equal(tt.output, floor)
	}
}

func Test_NoMath(t *testing.T) {

	suite.Run(t, new(TestSuite))
}
