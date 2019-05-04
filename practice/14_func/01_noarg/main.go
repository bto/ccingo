package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/14_func/01_noarg/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
