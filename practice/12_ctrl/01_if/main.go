package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/12_ctrl/01_if/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
