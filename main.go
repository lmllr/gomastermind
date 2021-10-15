package main

import (
	"fmt"

	"github.com/lmllr/gomastermind/packages/get"
)

func main() {
	code := []rune{49, 50, 51, 52}

	fmt.Println("🤖 Please give the machine some words/characters to work with:")
	ui := get.UsrInpt()
	fmt.Println(ui)

	fmt.Println(get.Fdbk(code, ui))
}
