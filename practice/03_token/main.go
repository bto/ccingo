package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/bto/ccingo/practice/03_token/cc"
)

func errorToken(tk cc.Token) {
	log.Fatal("予期しないトークンです: ", string(tk.Input))
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := cc.Tokenize(rd)

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	if tks[0].Ty != cc.TK_NUM {
		errorToken(tks[0])
	}
	fmt.Println("  mov rax,", tks[0].Val)

	for i := 1; tks[i].Ty != cc.TK_EOF; i++ {
		if tks[i].Ty == '+' {
			i++
			if tks[i].Ty != cc.TK_NUM {
				errorToken(tks[i])
			}
			fmt.Println("  add rax,", tks[i].Val)
			continue
		}

		if tks[i].Ty == '-' {
			i++
			if tks[i].Ty != cc.TK_NUM {
				errorToken(tks[i])
			}
			fmt.Println("  sub rax,", tks[i].Val)
			continue
		}

		errorToken(tks[i])
	}

	fmt.Println("  ret")
}
