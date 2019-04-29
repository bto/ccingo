package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/07_var/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	nds := cc.Parse(tks)
	cc.PrintAsm(nds)
}