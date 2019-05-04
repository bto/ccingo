package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/13_block/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
