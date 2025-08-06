package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello   world    ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   My BIG NAME IS    monster hunter!  ",
			expected: []string{"my", "big", "name", "is", "monster", "hunter!"},
		},
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of output should be: %d but it is: %d", len(c.expected), len(actual))
			break
		}
		for i, w := range actual {
			if c.expected[i] != w {
				t.Errorf("%s is not equal to %s", c.expected[i], w)
			}
		}
	}
}
