package finding

import "testing"

func TestBoyerMooreSearch(t *testing.T) {
	tests := []struct {
		txt, pattern string
		res          int
	}{
		{"", "", 0},
		{" ", "", 0},
		{"1", "", 0},
		{"a", "", 0},
		{"A", "", 0},
		{"_", "", 0},
		{"0 ", "", 0},
		{"01", "", 0},
		{"0a", "", 0},
		{"0A", "", 0},
		{"0_", "", 0},
		{"0 1", "", 0},
		{"0 a", "", 0},
		{"0 A", "", 0},
		{"0 _", "", 0},
		{"0a ", "", 0},
		{"0a1", "", 0},
		{"0aA", "", 0},
		{"0a_", "", 0},
		{"0A ", "", 0},
		{"0A1", "", 0},
		{"0Aa", "", 0},
		{"0A_", "", 0},
		{"0aA ", "", 0},
		{"0aA1", "", 0},
		{"0aA_", "", 0},
		{"0aA_1", "", 0},
		{"", "a", -1},
		{" ", "a", -1},
		{"1", "a", -1},
		{"a", "a", 0},
		{"A", "a", -1},
		{"_", "a", -1},
		{"0 ", "a", -1},
		{"01", "a", -1},
		{"0a", "a", 1},
		{"0A", "a", -1},
		{"0_", "a", -1},
		{"0 1", "a", -1},
		{"0 a", "a", 2},
		{"0 A", "a", -1},
		{"0 _", "a", -1},
		{"0a ", "a", 1},
		{"0a1", "a", 1},
		{"0aA", "a", 1},
		{"0a_", "a", 1},
		{"0A ", "a", -1},
		{"0A1", "a", -1},
		{"0Aa", "a", 2},
		{"0A_", "a", -1},
		{"0aA ", "a", 1},
		{"0aA1", "a", 1},
		{"0aA_", "a", 1},
		{"0aA_1", "a", 1},
		{"", "1a", -1},
		{" ", "1a", -1},
		{"1", "1a", -1},
		{"a", "1a", -1},
		{"A", "1a", -1},
		{"_", "1a", -1},
		{"0 ", "1a", -1},
		{"01", "1a", -1},
		{"0a", "1a", -1},
		{"0A", "1a", -1},
		{"0_", "1a", -1},
		{"0 1", "1a", -1},
		{"0 a", "1a", -1},
		{"0 A", "1a", -1},
		{"0 _", "1a", -1},
		{"0a ", "1a", -1},
		{"0a1", "1a", -1},
		{"0aA", "1a", -1},
		{"0a_", "1a", -1},
		{"0A ", "1a", -1},
		{"0A1", "1a", -1},
		{"0Aa", "1a", -1},
		{"0A_", "1a", -1},
		{"0aA ", "1a", -1},
		{"0aA1", "1a", -1},
		{"0aA_", "1a", -1},
		{"0aA_1", "1a", -1},
		{"0aA_1", "A_", 2},
		{"0aA_A_1", "A_", 2},
	}

	for _, test := range tests {
		if res := BoyerMooreSearch(test.txt, test.pattern); res != test.res {
			t.Fatalf("BoyerMooreSearch(%#v,%#v)\ngot %d,want %d", test.txt, test.pattern, res, test.res)
		}
	}
}
