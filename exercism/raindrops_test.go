package exercism

import "testing"

func TestConvert(t *testing.T) {
	for _, tc := range raindropsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Convert(tc.input); actual != tc.expected {
				t.Fatalf("Convert(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	for b.Loop() {
		for _, test := range raindropsTestCases {
			Convert(test.input)
		}
	}
}
