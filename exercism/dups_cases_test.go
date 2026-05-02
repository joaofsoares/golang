package exercism

type dupsTestCase struct {
	description string
	input       string
	expected    string
}

var dupsTestCases = []dupsTestCase{
	{
		description: "abbc",
		input:       "abbc",
		expected:    "ac",
	},
	{
		description: "abbbc",
		input:       "abbbc",
		expected:    "ac",
	},
	{
		description: "aaaabbbccd",
		input:       "aaaabbbccd",
		expected:    "d",
	},
	{
		description: "abbac",
		input:       "abbac",
		expected:    "c",
	},
	{
		description: "abcde",
		input:       "abcde",
		expected:    "abcde",
	},
	{
		description: "a",
		input:       "a",
		expected:    "a",
	},
	{
		description: "aa",
		input:       "aa",
		expected:    "",
	},
	{
		description: "ab",
		input:       "ab",
		expected:    "ab",
	},
	{
		description: "abc",
		input:       "abc",
		expected:    "abc",
	},
}
