package get

import (
	"bufio"
	"os"
)

func UsrInpt() []rune {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sc := []rune(scanner.Text())
	return sc
}
