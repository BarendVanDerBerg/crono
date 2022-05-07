package parser

import "testing"

var argumentTestCases = []struct {
	Name       string
	Arguments  []string
	ShouldFail bool
}{
	{
		Name:      "working case - simple command",
		Arguments: []string{"crono", "*/15 0 1,15 * 1-5 /usr/bin/find"},
	},
	{
		Name:      "working case - multi string command",
		Arguments: []string{"crono", "*/15 0 1,15 * 1-5 /usr/bin/find -r password.txt"},
	},
	{
		Name:       "breaking case - no command",
		Arguments:  []string{"crono", "*/15 0 1,15 * 1-5"},
		ShouldFail: true,
	},
	{
		Name:       "breaking case - missing schedule",
		Arguments:  []string{"crono", "*/15 0 1,15 * /usr/bin/find"},
		ShouldFail: true,
	},
}

func TestParseArguments(t *testing.T) {
	for _, test := range argumentTestCases {
		t.Run(test.Name, func(t *testing.T) {
			_, err := ParseArguments(test.Arguments)
			if err != nil && !test.ShouldFail {
				t.Error(err)
			}
		})
	}
}

func BenchmarkParseArguments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseArguments(argumentTestCases[0].Arguments)
	}
}
