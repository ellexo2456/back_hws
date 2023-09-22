package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tr = true
var fls = false
var zero = 0
var one = 1
var two = 2
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

var allFlagsDown Options = Options{C: fls, D: fls, U: fls, I: fls, F: zero, S: zero}

var goodCases = map[string]struct {
	input   []string
	result  []string
	options Options
}{
	"without params": {
		input: inputStrings,
		result: []string{
			"I love music.",
			"\n",
			"I Love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
		options: allFlagsDown,
	},
	"all strings different": {
		input: []string{
			"I still fall for you like suns do for skies",
			"Cerulean pouring in from your eyes",
			"Just a hollow moon that you colorize",
			"So powerful, I feel so small but so alive",
			"Like watching the Earth rise",
		},
		result: []string{
			"I still fall for you like suns do for skies",
			"Cerulean pouring in from your eyes",
			"Just a hollow moon that you colorize",
			"So powerful, I feel so small but so alive",
			"Like watching the Earth rise",
		},
		options: allFlagsDown,
	},
	"all strings the same": {
		input: []string{
			"the same string",
			"the same string",
			"the same string",
			"the same string",
			"the same string",
			"the same string",
			"the same string",
		},
		result: []string{
			"the same string",
		},
		options: allFlagsDown,
	},
	"all strings blank": {
		input: []string{
			"\n",
			"\n",
			"\n",
			"\n",
			"\n",
			"\n",
			"\n",
		},
		result: []string{
			"\n",
		},
		options: allFlagsDown,
	},
	"empty strings": {
		input:   []string{},
		result:  []string{},
		options: allFlagsDown,
	},
	"c flag": {
		input: inputStrings,
		result: []string{
			"3 I love music.",
			"1 \n",
			"2 I Love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
		options: Options{C: tr, D: fls, U: fls, I: fls, F: zero, S: zero},
	},
	"d flag": {
		input: inputStrings,
		result: []string{
			"I love music.",
			"I Love music of Kartik.",
			"I love music of Kartik.",
		},
		options: Options{C: fls, D: tr, U: fls, I: fls, F: zero, S: zero},
	},
	"u flag": {
		input: inputStrings,
		result: []string{
			"\n",
			"Thanks.",
		},
		options: Options{C: fls, D: fls, U: tr, I: fls, F: zero, S: zero},
	},
	"i flag": {
		input: []string{
			"I loVE music.",
			"i love muSic.",
			"I LOVE music.",
			"\n",
			"I LoVe mUsIc Of KaRtIk.",
			"i love music of kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		result: []string{
			"I loVE music.",
			"\n",
			"I LoVe mUsIc Of KaRtIk.",
			"Thanks.",
			"I love music of Kartik.",
		},
		options: Options{C: fls, D: fls, U: fls, I: tr, F: zero, S: zero},
	},
	"f flag": {
		input: []string{
			"I love music.",
			"We`re love music.",
			"and love music.",
			"\n",
			"I Love music of Kartik.",
			"I Love music of Kartik.",
			"Thanks.",
			"some love music of Kartik.",
			"I love music of Kartik.",
		},
		result: []string{
			"I love music.",
			"\n",
			"I Love music of Kartik.",
			"Thanks.",
			"some love music of Kartik.",
		},
		options: Options{C: fls, D: fls, U: fls, I: fls, F: one, S: zero},
	},
	"s flag": {
		input: []string{
			"I love music.",
			"We love music.",
			"and love music.",
			"\n",
			"He Love music of Kartik.",
			"We Love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		result: []string{
			"I love music.",
			"We love music.",
			"and love music.",
			"\n",
			"He Love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
		options: Options{C: fls, D: fls, U: fls, I: fls, F: zero, S: two},
	},
	"f and s flags": {
		input: []string{
			"I LOve music.",
			"We`re lOve music.",
			"and Love music.",
			"\n",
			"I Love music of Kartik.",
			"I Love music of Kartik.",
			"Thanks.",
			"some loVe music of Kartik.",
			"I love music of Kartik.",
		},
		result: []string{
			"I LOve music.",
			"\n",
			"I Love music of Kartik.",
			"Thanks.",
			"some loVe music of Kartik.",
			"I love music of Kartik.",
		},
		options: Options{C: fls, D: fls, U: fls, I: fls, F: one, S: two},
	},
	"f flag a lot": {
		input: []string{
			"I love music.",
			"We`re love music.",
			"and love music.",
			"\n",
			"I Love music of Kartik.",
			"I Love music of Kartik.",
			"Thanks.",
			"some love music of Kartik.",
			"I love music of Kartik.",
		},
		result: []string{
			"I love music.",
		},
		options: Options{C: fls, D: fls, U: fls, I: fls, F: 10, S: zero},
	},
	"s flag with rune strings": {
		input: []string{
			"Мне нравится музыка.",
			"Им нравится музыка.",
			"те нравится музыка.",
			"\n",
			"π ≠ 3.14",
			"å ≠ 3.14",
			"Пасиба.",
			"Мне нравится музыка.",
			"Мне нравится музыка.",
		},
		result: []string{
			"Мне нравится музыка.",
			"Им нравится музыка.",
			"\n",
			"π ≠ 3.14",
			"Пасиба.",
			"Мне нравится музыка.",
		},
		options: Options{C: fls, D: fls, U: fls, I: fls, F: zero, S: two},
	},
}

var badCases = map[string]struct {
	input   []string
	result  []string
	options Options
}{
	"undefined options": {
		input:   inputStrings,
		result:  nil,
		options: Options{},
	},
	"undefined strings": {
		input:   nil,
		result:  nil,
		options: allFlagsDown,
	},
	"flags c, d, u together": {
		input:   inputStrings,
		result:  nil,
		options: Options{C: tr, D: tr, U: tr, I: fls, F: zero, S: zero},
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
