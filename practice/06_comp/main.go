package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/06_comp/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)
	nd := tks.Parse()
	nd.PrintAsm()
}
