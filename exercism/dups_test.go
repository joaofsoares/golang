package exercism

import "testing"

func TestRemoveDups(t *testing.T) {
	for _, testCase := range dupsTestCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := Remove(testCase.input)
			if actual != testCase.expected {
				t.Errorf("got %s, expected %s", actual, testCase.expected)
			}
		})
	}
}
