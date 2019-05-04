package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/12_ctrl/03_for/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
