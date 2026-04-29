package exercism

import "testing"

func TestValid(t *testing.T) {

	for _, tc := range luhnTestCases {

		t.Run(tc.description, func(t *testing.T) {

			if actual := Valid(tc.input); actual != tc.expected {

				t.Fatalf("Valid(%q) = %t, want: %t", tc.input, actual, tc.expected)

			}

		})

	}

}

func BenchmarkValid(b *testing.B) {

	for b.Loop() {

		for _, tc := range luhnTestCases {

			Valid(tc.input)

		}

	}

}
