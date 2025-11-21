package unixsort_test

import (
	"os"
	"strings"
	"testing"

	"github.com/andreyxaxa/unixsort/pkg/unixsort"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSortFromFiles(t *testing.T) {
	inputB, err := os.ReadFile("../../pkg/unixsort/testdata/basicinput.txt")
	require.NoError(t, err)

	expectedB, err := os.ReadFile("../../pkg/unixsort/testdata/basicexpected.txt")
	require.NoError(t, err)

	cleanInput := strings.ReplaceAll(string(inputB), "\r", "")
	cleanExpected := strings.ReplaceAll(string(expectedB), "\r", "")

	input := strings.Split(strings.TrimRight(cleanInput, "\n"), "\n")
	expected := strings.Split(strings.TrimRight(cleanExpected, "\n"), "\n")

	p := unixsort.NewParams()
	actual := p.Sort(input)

	assert.Equal(t, expected, actual)
}

func TestSort(t *testing.T) {

	testCases := []struct {
		name     string
		params   *unixsort.Params
		input    []string
		expected []string
	}{
		{
			name:   "basic sort",
			params: &unixsort.Params{},
			input: []string{
				"There 9 was only half a worm in the apple.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
				"There 4 was only half a worm in the apple.",
				"There 3 was only half a worm in the apple.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"It ended up being much worse than that.",
				"It was supposed to be a dream vacation.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
			},
			expected: []string{
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"It ended up being much worse than that.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
				"It was supposed to be a dream vacation.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"There 3 was only half a worm in the apple.",
				"There 4 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
			},
		},
		{
			name:   "numeric",
			params: &unixsort.Params{Numeric: true},
			input: []string{
				"33",
				"11",
				"190",
				"12",
				"2891",
				"129",
				"9900",
				"1553",
				"124",
				"9",
				"1",
				"17",
				"2",
				"1984",
				"9901",
				"9899",
				"6",
				"17000",
				"22",
				"19",
			},
			expected: []string{
				"1",
				"2",
				"6",
				"9",
				"11",
				"12",
				"17",
				"19",
				"22",
				"33",
				"124",
				"129",
				"190",
				"1553",
				"1984",
				"2891",
				"9899",
				"9900",
				"9901",
				"17000",
			},
		},
		{
			name:   "reverse",
			params: &unixsort.Params{Reverse: true},
			input: []string{
				"There 9 was only half a worm in the apple.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
				"There 4 was only half a worm in the apple.",
				"There 3 was only half a worm in the apple.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"It ended up being much worse than that.",
				"It was supposed to be a dream vacation.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
			},
			expected: []string{
				"she 22 wondered. And then it dawned on her.",
				"Why 1 would only half a worm be living in an apple?",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"There 9 was only half a worm in the apple.",
				"There 4 was only half a worm in the apple.",
				"There 3 was only half a worm in the apple.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"It was supposed to be a dream vacation.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
				"It ended up being much worse than that.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
			},
		},
		{
			name:   "unique",
			params: &unixsort.Params{Unique: true},
			input: []string{
				"There 9 was only half a worm in the apple.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
				"There 4 was only half a worm in the apple.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"There 3 was only half a worm in the apple.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"It ended up being much worse than that.",
				"There 4 was only half a worm in the apple.",
				"There 4 was only half a worm in the apple.",
				"It was supposed to be a dream vacation.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
			},
			expected: []string{
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"It ended up being much worse than that.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
				"It was supposed to be a dream vacation.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"There 3 was only half a worm in the apple.",
				"There 4 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
			},
		},
		{
			name:   "column",
			params: &unixsort.Params{Column: 2},
			input: []string{
				"There 9 was only half a worm in the apple.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
				"There 4 was only half a worm in the apple.",
				"There 3 was only half a worm in the apple.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"It ended up being much worse than that.",
				"It was supposed to be a dream vacation.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
			},
			expected: []string{
				"Why 1 would only half a worm be living in an apple?",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"she 22 wondered. And then it dawned on her.",
				"There 3 was only half a worm in the apple.",
				"There 4 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"It ended up being much worse than that.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
				"It was supposed to be a dream vacation.",
			},
		},
		{
			name:   "human",
			params: &unixsort.Params{Human: true},
			input: []string{
				"12M",
				"224G",
				"1T",
				"22M",
				"1K",
				"121351215K",
				"1234T",
				"1M",
				"983561G",
				"1024K",
				"111M",
				"204K",
				"2225K",
				"10964M",
				"11T",
				"3T",
				"4M",
				"1246K",
			},
			expected: []string{
				"1K",
				"204K",
				"1M",
				"1024K",
				"1246K",
				"2225K",
				"4M",
				"12M",
				"22M",
				"111M",
				"10964M",
				"121351215K",
				"224G",
				"1T",
				"3T",
				"11T",
				"983561G",
				"1234T",
			},
		},
		{
			name:   "months",
			params: &unixsort.Params{Month: true},
			input: []string{
				"There 9 was Jul only half a worm in the apple.",
				"At 12 first, Judy didn't Jan quite comprehend what this meant.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And Dec then it dawned on her.",
				"There 4 was only half a aug worm in the apple.",
			},
			expected: []string{
				"At 12 first, Judy didn't Jan quite comprehend what this meant.",
				"There 4 was only half a aug worm in the apple.",
				"There 9 was Jul only half a worm in the apple.",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And Dec then it dawned on her.",
			},
		},
		{
			name:   "blanks ignore",
			params: &unixsort.Params{BlanksIgnore: true},
			input: []string{
				"There 9 was only half a worm in the apple       ",
				"There 2 was only half a worm in the apple  ",
				"There 1 was only half a worm in the apple",
				"There 4 was only half a worm in the apple",
			},
			expected: []string{
				"There 1 was only half a worm in the apple",
				"There 2 was only half a worm in the apple  ",
				"There 4 was only half a worm in the apple",
				"There 9 was only half a worm in the apple       ",
			},
		},
		{
			name:   "-nruk 2",
			params: &unixsort.Params{Numeric: true, Reverse: true, Unique: true, Column: 2},
			input: []string{
				"There 9 was only half a worm in the apple.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"Why 1 would only half a worm be living in an apple?",
				"Why 1 would only half a worm be living in an apple?",
				"Why 1 would only half a worm be living in an apple?",
				"Why 1 would only half a worm be living in an apple?",
				"Why 1 would only half a worm be living in an apple?",
				"she 22 wondered. And then it dawned on her.",
				"There 4 was only half a worm in the apple.",
				"There 3 was only half a worm in the apple.",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"It ended up being much worse than that",
				"There 9 was only half a worm in the apple.",
				"It was supposed to be a dream vacation.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
				"It had been the light at the end of both their tunnels.",
				"Now that the dream vacation was only a week away, the virus had stopped all air travel.",
				"He sat staring at the person in the train stopped at the station going in the opposite direction.",
				"She 22 sat staring ahead, never noticing that she was being watched.",
				"Both 11 trains began to move and he knew that in another timeline or in another universe,",
				"There 7 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"There 9 was only half a worm in the apple.",
				"There was only half a worm in the apple.",
				"they 4 had been happy together.",
			},
			expected: []string{
				"It was supposed to be a dream vacation.",
				"There was only half a worm in the apple.",
				"Now that the dream vacation was only a week away, the virus had stopped all air travel.",
				"He sat staring at the person in the train stopped at the station going in the opposite direction.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It had been what they had been looking forward to through all the turmoil and negativity around them.",
				"It had been the light at the end of both their tunnels.",
				"They had planned it over a year in advance so that it would be perfect in every way.",
				"It ended up being much worse than that",
				"She 22 sat staring ahead, never noticing that she was being watched.",
				"she 22 wondered. And then it dawned on her.",
				"At 12 first, Judy didn't quite comprehend what this meant.",
				"Both 11 trains began to move and he knew that in another timeline or in another universe,",
				"Judy 10 quickly spit out the bite she had just taken expecting to see the other half of the worm.",
				"There 9 was only half a worm in the apple.",
				"There 7 was only half a worm in the apple.",
				"There 4 was only half a worm in the apple.",
				"they 4 had been happy together.",
				"There 3 was only half a worm in the apple.",
				"Why 1 would only half a worm be living in an apple?",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.params.Sort(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
