package models

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		slice       []string
		str         string
		expected    bool
		description string
	}{
		{[]string{"includes", "compliances"}, "includes", true, "correctly includes"},
		{[]string{"includes", "compliances"}, "include", false, "incorrect spelling"},
		{[]string{}, "includes", false, "empty slice"},
	}

	for _, test := range tests {
		value := contains(test.slice, test.str)
		if value != test.expected {
			t.Errorf("\n returned %v expected %v, %v", value, test.expected, test.description)
		}
	}
}
