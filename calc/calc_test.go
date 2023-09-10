package calc

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

var inputStrings = []string{
	"I love music.",
	"I love music.",
	"I love music.",
	"\n",
	"I Love music of Kartik.",
	"I Love music of Kartik.",
	"Thanks.",
	"I love music of Kartik.",
	"I love music of Kartik.",
}

var goodCases = map[string]struct {
	input  string
	result float64
}{
	"common": {
		input: "(10*3)-((12/2)+(4*2))+((15-6)/3)",
		result: []string{
			"I love music.",
			"\n",
			"I Love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
}

var badCases = map[string]struct {
	input   []string
	result  []string
	options Options
	err     error
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
