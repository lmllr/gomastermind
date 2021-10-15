package get

import (
	"testing"

	"github.com/lmllr/gomastermind/packages/get"
)

func TestFeedback(t *testing.T) {
	codemaker := []rune{49, 50, 51, 52}
	codebreaker := []rune{'🤪', '😛', '😜', '😝'}

	want := []string{".", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q", got, want)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{49, '😛', 51, '😝'}

	want = []string{"*", "*", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q", got, want)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{52, 51, 50, 49}

	want = []string{"+", "+", "+", "+"}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q", got, want)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{23, 50, 50, 50}

	want = []string{"*", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q", got, want)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{23, 49, 49, 49}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q", got, want)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{23, 49, 51, 49}

	want = []string{"*", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q", got, want)
	}
}

// helper func to compare 2 slices
// check if a and b contain the same elements
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
