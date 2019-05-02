package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/11_ctrl/02_while/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	nds := tks.Parse()
	cc.PrintAsm(nds)
}
