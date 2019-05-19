package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintLlvm()
}
