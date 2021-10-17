package main

import (
	"fmt"

	"github.com/lmllr/gomastermind/packages/get"
)

func main() {
	code := []rune{49, 50, 51, 52}

	fmt.Println("🤖 Please give the machine some words/characters to work with:")

	for i := 0; i < 12; i++ {
		ui := get.UsrInpt()
		fmt.Println(ui)

		fdbk := get.Fdbk(code, ui)
		fmt.Println(fdbk)

		if get.Compare(code, ui) {
			fmt.Println("CRACKED 🔓")
			os.Exit(0)
		}
	}
}
