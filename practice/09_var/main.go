package main

import (
	"bufio"
	"os"

	"github.com/bto/ccingo/practice/09_var/cc"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	cc.Tokenize(rd).Parse().PrintAsm()
}
