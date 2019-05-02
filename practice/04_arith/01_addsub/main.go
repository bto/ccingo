package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/04_arith/01_addsub/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
