package get

import (
	"sort"
)

// check users guess on the code
// provides feedback by placing from zero to four chars to
// determine the correctness of position and color

// A "*" is placed for each char from the guess which
// is correct in both color and position.
// A "+" indicates the existence of a correct char placed in the wrong position.
func Fdbk(c []rune, g []rune) []string {
	b := map[int][2]bool{}

	for i, v := range c {

		if v == g[i] {
			b[i] = [2]bool{true, true}
		} else {
			for _, v2 := range g {
				if v2 == v {
					b[i] = [2]bool{false, true}
					break
				} else {
					b[i] = [2]bool{false, false}
				}
			}
		}

	}

	// fmt.Println(b)

	str := []string{}
	for _, v := range b {
		if v[0] && v[1] {
			str = append(str, "*")
		} else if v[1] {
			str = append(str, "+")
		} else {
			str = append(str, ".")
		}
	}
	sort.Strings(str)
	return str
}
