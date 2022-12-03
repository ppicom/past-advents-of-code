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
			input:  "",
			output: 0,
		},
	}

	for _, tt := range tests {

		suite.NotEqual(tt.input, tt.output)
	}

}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
