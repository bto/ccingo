package cc

import (
	"fmt"
	"log"
)

func errorToken(tk token) {
	log.Fatal("予期しないトークンです: ", string(tk.input))
}

func PrintAsm(tks []token) {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	if tks[0].ty != TK_NUM {
		errorToken(tks[0])
	}
	fmt.Println("  mov rax,", tks[0].val)

	for i := 1; tks[i].ty != TK_EOF; i++ {
		if tks[i].ty == '+' {
			i++
			if tks[i].ty != TK_NUM {
				errorToken(tks[i])
			}
			fmt.Println("  add rax,", tks[i].val)
			continue
		}

		if tks[i].ty == '-' {
			i++
			if tks[i].ty != TK_NUM {
				errorToken(tks[i])
			}
			fmt.Println("  sub rax,", tks[i].val)
			continue
		}

		errorToken(tks[i])
	}

	fmt.Println("  ret")
}
