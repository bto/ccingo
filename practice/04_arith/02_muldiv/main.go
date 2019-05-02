package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/04_arith/02_muldiv/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	nd := tks.Parse()
	cc.PrintAsm(nd)
}
