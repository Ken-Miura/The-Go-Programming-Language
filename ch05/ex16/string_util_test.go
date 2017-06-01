// Copyright 2017 Ken Mirua
package ex16

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		s3       string
		sep      string
		expected string
	}{
		{"a", "b", "c", ", ", strings.Join([]string{"a", "b", "c"}, ", ")},
		{"a", "b", "c", "", strings.Join([]string{"a", "b", "c"}, "")},
		{"", "", "", "", strings.Join([]string{"", "", ""}, "")},
		{"a", "b", "c", "", strings.Join([]string{"a", "b", "c"}, "")},
	}

	for _, test := range tests {
		result := Join(test.sep, test.s1, test.s2, test.s3)
		if result != test.expected {
			t.Fatalf("%s is expected but result is %s", test.expected, result)
		}
	}
}
