package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/14_func_call/02_args/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
