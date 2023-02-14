package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertTime(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{

		{name: "Test 1 Night", input: "12:00:00AM", expected: "00:00:00"},
		{name: "Test 2 Night", input: "12:01:00AM", expected: "00:01:00"},
		{name: "Test 3 Night", input: "01:23:45AM", expected: "01:23:45"},
		{name: "Test 4 Day", input: "12:00:00PM", expected: "12:00:00"},
		{name: "Test 5 Day", input: "12:01:00PM", expected: "12:01:00"},
		{name: "Test 6 Day", input: "01:23:45PM", expected: "13:23:45"},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual, err := ConvertTime(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
