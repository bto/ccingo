package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/10_var/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	nds := tks.Parse()
	cc.PrintAsm(nds)
}
