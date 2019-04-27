package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/03_token/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	cc.PrintAsm(tks)
}
