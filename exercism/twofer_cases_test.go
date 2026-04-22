package exercism

type twoFerTestCase struct {
	description string

	input string

	expected string
}

var twoFerTestCases = []twoFerTestCase{

	{

		description: "no name given",

		input: "",

		expected: "One for you, one for me.",
	},

	{

		description: "a name given",

		input: "Alice",

		expected: "One for Alice, one for me.",
	},

	{

		description: "another name given",

		input: "Bob",

		expected: "One for Bob, one for me.",
	},
}
