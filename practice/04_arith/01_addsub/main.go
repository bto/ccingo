package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/04_arith/01_addsub/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	nd := tks.Parse()
	nd.PrintAsm()
}
