package exercism

import (
	"testing"
)

func TestScore(t *testing.T) {
	for _, tc := range dartsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Score(tc.x, tc.y)
			if actual != tc.expected {
				t.Errorf("Score(%#v, %#v) = %#v, want: %#v", tc.x, tc.y, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	for b.Loop() {
		for _, tc := range dartsTestCases {
			Score(tc.x, tc.y)
		}
	}
}
