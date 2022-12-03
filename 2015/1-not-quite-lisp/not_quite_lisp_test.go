package notquitelisp

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {}

func (suite *TestSuite) TestXxx() {

	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "(",
			output: 1,
		},
		{
			input:  ")",
			output: -1,
		},
		{
			input:  "()",
			output: 0,
		},
		{
			input:  "(())",
			output: 0,
		},
		{
			input:  "()()",
			output: 0,
		},
		{
			input:  "(((",
			output: 3,
		},
		{
			input:  "(()(()(",
			output: 3,
		},
		{
			input:  "))(((((",
			output: 3,
		},
		{
			input:  "())",
			output: -1,
		},
		{
			input:  "))(",
			output: -1,
		},
		{
			input:  ")))",
			output: -3,
		},
		{
			input:  ")())())",
			output: -3,
		},
	}

	for _, tt := range tests {

		floor := TranslateDirectionsToFloor(tt.input)
		suite.Equal(tt.output, floor)
	}
}

func TestShouldTest(t *testing.T) {

	suite.Run(t, new(TestSuite))
}
