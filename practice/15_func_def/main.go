package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/15_func_def/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
