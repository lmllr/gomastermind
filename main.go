package main

import (
	"fmt"

	"github.com/lmllr/gomastermind/packages/get"
)

func main() {
	fmt.Println("🤖 Please give the machine some words/characters to work with:")
	ui := get.UsrInpt()
	fmt.Println(ui)
}
