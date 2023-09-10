package calc

import (
	"errors"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

var goodCases = map[string]struct {
	input  string
	result float64
}{
	"common": {
		input:  "(10*3)-((12/2)+(4*2))+((15-6)/3)",
		result: 19,
	},
	"with whitespaces": {
		input:  "((18 + 6) * 5) - (48 / (3 - 1)) + (20 - (2 * 2))",
		result: 112,
	},
	"with negatives inside": {
		input:  "(10*3)-((-12/2)+(4*(-2)))+((15-6)/3)",
		result: 47,
	},
	"with two minuses": {
		input:  "(3*2/2)+((9*3)-(14-(-6)))*(5-2)",
		result: 24,
	},
	"without parenthesis": {
		input:  "3*2/2+9*3-14-6*5-2",
		result: -16,
	},
	"only one number": {
		input:  "33456",
		result: 33456,
	},
	"with negative forehead": {
		input:  "-336+232*24",
		result: 5232,
	},
	"division by zero": {
		input:  "-336+232*/0",
		result: math.Inf(1),
	},
	"multiply on 1": {
		input:  "56789*1",
		result: math.Inf(56789),
	},
}

var badCases = map[string]struct {
	input  string
	result float64
}{
	"undefined options": {
		input:   inputStrings,
		result:  nil,
		options: Options{},
		err:     errors.New("Empty options"),
	},
	"undefined strings": {
		input:   nil,
		result:  nil,
		options: allFlagsDown,
		err:     errors.New("Empty input"),
	},
	"flags c, d, u together": {
		input:   inputStrings,
		result:  nil,
		options: Options{C: &tr, D: &tr, U: &tr, I: &fls, F: &zero, S: &zero},
		err:     errors.New("You`re can`t use flags c,d and u together"),
	},
}

func TestGoodCases(t *testing.T) {
	for name, test := range goodCases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Unique(test.input, test.options)
			expected := test.result

			require.Equal(t, expected, got)
			require.Equal(t, nil, err)
		})
	}
}

func TestBadCases(t *testing.T) {
	for name, test := range badCases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Unique(test.input, test.options)
			expected := test.result

			require.Equal(t, expected, got)
			require.NotEqual(t, nil, err)
		})
	}
}
