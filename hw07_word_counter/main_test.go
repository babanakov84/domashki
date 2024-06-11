package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "hey, teacher",
			expected: map[string]int{
				"hey":     1,
				"teacher": 1,
			},
		},
		{
			input: "We do not need no education, We do not need no thought control.",
			expected: map[string]int{
				"we":        2,
				"do":        2,
				"not":       2,
				"need":      2,
				"no":        2,
				"education": 1,
				"thought":   1,
				"control":   1,
			},
		},
		{
			input: "teachers, leave those kids alone",
			expected: map[string]int{
				"teachers": 1,
				"leave":    1,
				"those":    1,
				"kids":     1,
				"alone":    1,
			},
		},
		{
			input:    "  ",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := countWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для входной строки '%s' ожидался результат %v, но получен %v", test.input, test.expected, result)
		}
	}
}
