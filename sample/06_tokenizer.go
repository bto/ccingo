package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	TK_NUM = iota + 256
	TK_EOF
)

type token struct {
	ty, val int
	input []byte
}

func tokenize(rd *bufio.Reader) (tks []token) {
	var c byte
	var err error
	var tk token

	space := byte(' ')
	lf := byte('\n')
	plus := byte('+')
	minus := byte('-')
	zero := byte('0')
	nine := byte('9')

	for c, err = rd.ReadByte(); err == nil; {
		if c == space || c == lf {
			c, err = rd.ReadByte()
			continue
		}

		if c == plus || c == minus {
			tk := token {
				ty: int(c),
				input: []byte{c},
			}
			tks = append(tks, tk)

			c, err = rd.ReadByte()
			continue
		}

		if zero <= c && c <= nine {
			tk, c, err = tokenizeNum(rd, c)
			tks = append(tks, tk)
			continue
		}

		log.Fatal("トークナイズできません: ", string([]byte{c}))
	}
	if err != io.EOF {
		log.Fatal(err)
	}

	tk = token {
		ty: TK_EOF,
	}
	tks = append(tks, tk)

	return
}

func tokenizeNum(rd *bufio.Reader, v byte) (tk token, c byte, err error) {
	zero := byte('0')
	nine := byte('9')

	var num []byte
	for c = v; err == nil; c, err = rd.ReadByte() {
		if c < zero || nine < c {
			break
		}

		num = append(num, c)
	}

	val, err := strconv.Atoi(string(num))
	if err != nil {
		log.Fatal(err)
	}

	tk = token {
		ty: TK_NUM,
		val: val,
		input: num,
	}

	return
}

func errorToken(tk token) {
	log.Fatal("予期しないトークンです: ", string(tk.input))
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := tokenize(rd)

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
