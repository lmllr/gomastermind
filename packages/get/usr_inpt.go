package get

import (
	"bufio"
	"os"
)

// scan terminal for user input
// only valid inputs are four single characters
// examples:
//	1234
//	5314
//	1111
func UsrInpt() []rune {
	scanner := bufio.NewScanner(os.Stdin)

	var usrInpt []rune
	for {
		scanner.Scan()
		usrInpt = []rune(scanner.Text())

		if len(usrInpt) == 4 {
			break
		}
	}
	return usrInpt
}
