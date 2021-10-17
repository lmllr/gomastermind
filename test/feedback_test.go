package get

import (
	"fmt"
	"testing"

	"github.com/lmllr/gomastermind/packages/get"
)

func TestFeedback(t *testing.T) {
	codemaker := []rune{49, 50, 51, 52}
	codebreaker := []rune{'🤪', '😛', '😜', '😝'}

	want := []string{".", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{49, '😛', 51, '😝'}

	want = []string{"*", "*", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{52, 51, 50, 49}

	want = []string{"+", "+", "+", "+"}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{23, 50, 50, 50}

	want = []string{"*", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{23, 49, 49, 49}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{49, 50, 51, 52}
	codebreaker = []rune{23, 49, 51, 49}

	want = []string{"*", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{51, 50, 54, 49}
	codebreaker = []rune{49, 97, 97, 97}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{51, 54, 52, 52}
	codebreaker = []rune{52, 51, 51, 51}

	want = []string{"+", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{54, 51, 51, 51}
	codebreaker = []rune{51, 49, 49, 49}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	// ...
	codemaker = []rune{51, 52, 53, 49}
	codebreaker = []rune{49, 49, 49, 50}

	want = []string{"+", ".", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{51, 54, 54, 50}
	codebreaker = []rune{54, 54, 54, 54}

	want = []string{"*", "*", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{52, 54, 51, 51}
	codebreaker = []rune{51, 51, 49, 50}

	want = []string{"+", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{49, 49, 53, 53}
	codebreaker = []rune{53, 53, 49, 49}

	want = []string{"+", "+", "+", "+"}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
	}

	codemaker = []rune{50, 53, 50, 51}
	codebreaker = []rune{49, 50, 49, 50}

	want = []string{"+", "+", ".", "."}
	if got := get.Fdbk(codemaker, codebreaker); !equal(got, want) {
		t.Errorf("Fdbk() = %q, want %q\n\n", got, want)
		fmt.Printf("\tcode: %q,\n\tguess: %q\n\n", codemaker, codebreaker)
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
