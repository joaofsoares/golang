package exercism

import "testing"

func TestShareWith(t *testing.T) {

	for _, tc := range twoFerTestCases {

		t.Run(tc.description, func(t *testing.T) {

			got := ShareWith(tc.input)

			if got != tc.expected {

				t.Fatalf("ShareWith(%q)\n got: %q\nwant: %q", tc.input, got, tc.expected)

			}

		})

	}

}

func BenchmarkShareWith(b *testing.B) {

	for b.Loop() {

		for _, test := range twoFerTestCases {

			ShareWith(test.input)

		}

	}

}
